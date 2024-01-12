package repoupdater

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc"

	"github.com/sourcegraph/log/logtest"

	"github.com/sourcegraph/sourcegraph/cmd/repo-updater/internal/scheduler"
	"github.com/sourcegraph/sourcegraph/cmd/repo-updater/internal/syncer"
	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/conf/reposource"
	"github.com/sourcegraph/sourcegraph/internal/database"
	"github.com/sourcegraph/sourcegraph/internal/database/dbmocks"
	"github.com/sourcegraph/sourcegraph/internal/database/dbtest"
	"github.com/sourcegraph/sourcegraph/internal/extsvc"
	"github.com/sourcegraph/sourcegraph/internal/extsvc/awscodecommit"
	"github.com/sourcegraph/sourcegraph/internal/extsvc/github"
	"github.com/sourcegraph/sourcegraph/internal/extsvc/gitlab"
	"github.com/sourcegraph/sourcegraph/internal/gitserver"
	internalgrpc "github.com/sourcegraph/sourcegraph/internal/grpc"
	"github.com/sourcegraph/sourcegraph/internal/grpc/defaults"
	"github.com/sourcegraph/sourcegraph/internal/observation"
	"github.com/sourcegraph/sourcegraph/internal/repos"
	"github.com/sourcegraph/sourcegraph/internal/repoupdater"
	"github.com/sourcegraph/sourcegraph/internal/repoupdater/protocol"
	proto "github.com/sourcegraph/sourcegraph/internal/repoupdater/v1"
	"github.com/sourcegraph/sourcegraph/internal/timeutil"
	"github.com/sourcegraph/sourcegraph/internal/types"
	"github.com/sourcegraph/sourcegraph/internal/types/typestest"
	"github.com/sourcegraph/sourcegraph/lib/errors"
)

func TestServer_EnqueueRepoUpdate(t *testing.T) {
	ctx := context.Background()

	svc := types.ExternalService{
		Kind: extsvc.KindGitHub,
		Config: extsvc.NewUnencryptedConfig(`{
"url": "https://github.com",
"token": "secret-token",
"repos": ["owner/name"]
}`),
	}

	repo := types.Repo{
		ID:   1,
		Name: "github.com/foo/bar",
		ExternalRepo: api.ExternalRepoSpec{
			ID:          "bar",
			ServiceType: extsvc.TypeGitHub,
			ServiceID:   "https://github.com",
		},
		Metadata: new(github.Repository),
	}

	initStore := func(db database.DB) database.DB {
		if err := db.ExternalServices().Upsert(ctx, &svc); err != nil {
			t.Fatal(err)
		}
		if err := db.Repos().Create(ctx, &repo); err != nil {
			t.Fatal(err)
		}
		return db
	}

	type testCase struct {
		name string
		repo api.RepoName
		res  *protocol.RepoUpdateResponse
		err  string
		init func(database.DB) database.DB
	}

	testCases := []testCase{{
		name: "returns an error on store failure",
		init: func(realDB database.DB) database.DB {
			mockRepos := dbmocks.NewMockRepoStore()
			mockRepos.ListFunc.SetDefaultReturn(nil, errors.New("boom"))
			realStore := initStore(realDB)
			mockStore := dbmocks.NewMockDBFrom(realStore)
			mockStore.ReposFunc.SetDefaultReturn(mockRepos)
			return mockStore
		},
		err: `store.list-repos: boom`,
	}, {
		name: "missing repo",
		init: initStore,
		repo: "foo",
		err:  `repo foo not found with response: repo "foo" not found in store`,
	}, {
		name: "existing repo",
		repo: repo.Name,
		init: initStore,
		res: &protocol.RepoUpdateResponse{
			ID:   repo.ID,
			Name: string(repo.Name),
		},
	}}

	logger := logtest.Scoped(t)
	for _, tc := range testCases {
		tc := tc
		ctx := context.Background()

		t.Run(tc.name, func(t *testing.T) {
			sqlDB := dbtest.NewDB(t)
			store := tc.init(database.NewDB(logger, sqlDB))

			s := &Server{Logger: logger, DB: store, Scheduler: &fakeScheduler{}}
			gs := grpc.NewServer(defaults.ServerOptions(logger)...)
			proto.RegisterRepoUpdaterServiceServer(gs, s)

			srv := httptest.NewServer(internalgrpc.MultiplexHandlers(gs, http.NotFoundHandler()))
			defer srv.Close()

			cli := repoupdater.NewClient(srv.URL)
			if tc.err == "" {
				tc.err = "<nil>"
			}

			res, err := cli.EnqueueRepoUpdate(ctx, tc.repo)
			if have, want := fmt.Sprint(err), tc.err; !strings.Contains(have, want) {
				t.Errorf("have err: %q, want: %q", have, want)
			}

			if have, want := res, tc.res; !reflect.DeepEqual(have, want) {
				t.Errorf("response: %s", cmp.Diff(have, want))
			}
		})
	}
}

func TestServer_RepoLookup(t *testing.T) {
	db := database.NewDB(logtest.Scoped(t), dbtest.NewDB(t))

	ctx := context.Background()
	clock := timeutil.NewFakeClock(time.Now(), 0)
	now := clock.Now()

	githubSource := types.ExternalService{
		Kind:         extsvc.KindGitHub,
		CloudDefault: true,
		Config: extsvc.NewUnencryptedConfig(`{
"url": "https://github.com",
"token": "secret-token",
"repos": ["owner/name"]
}`),
	}
	awsSource := types.ExternalService{
		Kind: extsvc.KindAWSCodeCommit,
		Config: extsvc.NewUnencryptedConfig(`
{
  "region": "us-east-1",
  "accessKeyID": "abc",
  "secretAccessKey": "abc",
  "gitCredentials": {
    "username": "user",
    "password": "pass"
  }
}
`),
	}
	gitlabSource := types.ExternalService{
		Kind:         extsvc.KindGitLab,
		CloudDefault: true,
		Config: extsvc.NewUnencryptedConfig(`
{
  "url": "https://gitlab.com",
  "token": "abc",
  "projectQuery": ["none"]
}
`),
	}
	npmSource := types.ExternalService{
		Kind: extsvc.KindNpmPackages,
		Config: extsvc.NewUnencryptedConfig(`
{
  "registry": "npm.org"
}
`),
	}

	if err := db.ExternalServices().Upsert(ctx, &githubSource, &awsSource, &gitlabSource, &npmSource); err != nil {
		t.Fatal(err)
	}

	githubRepository := &types.Repo{
		Name:        "github.com/foo/bar",
		Description: "The description",
		Archived:    false,
		Fork:        false,
		CreatedAt:   now,
		UpdatedAt:   now,
		ExternalRepo: api.ExternalRepoSpec{
			ID:          "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==",
			ServiceType: extsvc.TypeGitHub,
			ServiceID:   "https://github.com/",
		},
		Sources: map[string]*types.SourceInfo{
			githubSource.URN(): {
				ID:       githubSource.URN(),
				CloneURL: "git@github.com:foo/bar.git",
			},
		},
		Metadata: &github.Repository{
			ID:            "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==",
			URL:           "github.com/foo/bar",
			DatabaseID:    1234,
			Description:   "The description",
			NameWithOwner: "foo/bar",
		},
	}

	awsCodeCommitRepository := &types.Repo{
		Name:        "git-codecommit.us-west-1.amazonaws.com/stripe-go",
		Description: "The stripe-go lib",
		Archived:    false,
		Fork:        false,
		CreatedAt:   now,
		ExternalRepo: api.ExternalRepoSpec{
			ID:          "f001337a-3450-46fd-b7d2-650c0EXAMPLE",
			ServiceType: extsvc.TypeAWSCodeCommit,
			ServiceID:   "arn:aws:codecommit:us-west-1:999999999999:",
		},
		Sources: map[string]*types.SourceInfo{
			awsSource.URN(): {
				ID:       awsSource.URN(),
				CloneURL: "git@git-codecommit.us-west-1.amazonaws.com/v1/repos/stripe-go",
			},
		},
		Metadata: &awscodecommit.Repository{
			ARN:          "arn:aws:codecommit:us-west-1:999999999999:stripe-go",
			AccountID:    "999999999999",
			ID:           "f001337a-3450-46fd-b7d2-650c0EXAMPLE",
			Name:         "stripe-go",
			Description:  "The stripe-go lib",
			HTTPCloneURL: "https://git-codecommit.us-west-1.amazonaws.com/v1/repos/stripe-go",
			LastModified: &now,
		},
	}

	gitlabRepository := &types.Repo{
		Name:        "gitlab.com/gitlab-org/gitaly",
		Description: "Gitaly is a Git RPC service for handling all the git calls made by GitLab",
		URI:         "gitlab.com/gitlab-org/gitaly",
		CreatedAt:   now,
		UpdatedAt:   now,
		ExternalRepo: api.ExternalRepoSpec{
			ID:          "2009901",
			ServiceType: extsvc.TypeGitLab,
			ServiceID:   "https://gitlab.com/",
		},
		Sources: map[string]*types.SourceInfo{
			gitlabSource.URN(): {
				ID:       gitlabSource.URN(),
				CloneURL: "https://gitlab.com/gitlab-org/gitaly.git",
			},
		},
		Metadata: &gitlab.Project{
			ProjectCommon: gitlab.ProjectCommon{
				ID:                2009901,
				PathWithNamespace: "gitlab-org/gitaly",
				Description:       "Gitaly is a Git RPC service for handling all the git calls made by GitLab",
				WebURL:            "https://gitlab.com/gitlab-org/gitaly",
				HTTPURLToRepo:     "https://gitlab.com/gitlab-org/gitaly.git",
				SSHURLToRepo:      "git@gitlab.com:gitlab-org/gitaly.git",
			},
			Visibility: "",
			Archived:   false,
		},
	}

	npmRepository := &types.Repo{
		Name: "npm/package",
		URI:  "npm/package",
		ExternalRepo: api.ExternalRepoSpec{
			ID:          "npm/package",
			ServiceType: extsvc.TypeNpmPackages,
			ServiceID:   extsvc.TypeNpmPackages,
		},
		Sources: map[string]*types.SourceInfo{
			npmSource.URN(): {
				ID:       npmSource.URN(),
				CloneURL: "npm/package",
			},
		},
		Metadata: &reposource.NpmMetadata{Package: func() *reposource.NpmPackageName {
			p, _ := reposource.NewNpmPackageName("", "package")
			return p
		}()},
	}

	testCases := []struct {
		name        string
		args        protocol.RepoLookupArgs
		stored      types.Repos
		result      *protocol.RepoLookupResult
		src         repos.Source
		assert      typestest.ReposAssertion
		assertDelay time.Duration
		err         string
	}{
		{
			name: "found - aws code commit",
			args: protocol.RepoLookupArgs{
				Repo: api.RepoName("git-codecommit.us-west-1.amazonaws.com/stripe-go"),
			},
			stored: []*types.Repo{awsCodeCommitRepository},
			result: &protocol.RepoLookupResult{Repo: &protocol.RepoInfo{
				ExternalRepo: api.ExternalRepoSpec{
					ID:          "f001337a-3450-46fd-b7d2-650c0EXAMPLE",
					ServiceType: extsvc.TypeAWSCodeCommit,
					ServiceID:   "arn:aws:codecommit:us-west-1:999999999999:",
				},
				Name:        "git-codecommit.us-west-1.amazonaws.com/stripe-go",
				Description: "The stripe-go lib",
				VCS:         protocol.VCSInfo{URL: "git@git-codecommit.us-west-1.amazonaws.com/v1/repos/stripe-go"},
				Links: &protocol.RepoLinks{
					Root:   "https://us-west-1.console.aws.amazon.com/codesuite/codecommit/repositories/stripe-go/browse",
					Tree:   "https://us-west-1.console.aws.amazon.com/codesuite/codecommit/repositories/stripe-go/browse/{rev}/--/{path}",
					Blob:   "https://us-west-1.console.aws.amazon.com/codesuite/codecommit/repositories/stripe-go/browse/{rev}/--/{path}",
					Commit: "https://us-west-1.console.aws.amazon.com/codesuite/codecommit/repositories/stripe-go/commit/{commit}",
				},
			}},
		},
		{
			name: "not synced from non public codehost",
			args: protocol.RepoLookupArgs{
				Repo: api.RepoName("github.private.corp/a/b"),
			},
			src:    repos.NewFakeSource(&githubSource, nil),
			result: &protocol.RepoLookupResult{ErrorNotFound: true},
			err:    fmt.Sprintf("repository not found (name=%s notfound=%v)", api.RepoName("github.private.corp/a/b"), true),
		},
		{
			name: "synced - npm package host",
			args: protocol.RepoLookupArgs{
				Repo: api.RepoName("npm/package"),
			},
			stored: []*types.Repo{},
			src:    repos.NewFakeSource(&npmSource, nil, npmRepository),
			result: &protocol.RepoLookupResult{Repo: &protocol.RepoInfo{
				ExternalRepo: npmRepository.ExternalRepo,
				Name:         npmRepository.Name,
				VCS:          protocol.VCSInfo{URL: string(npmRepository.Name)},
			}},
			assert: typestest.AssertReposEqual(npmRepository),
		},
		{
			name: "synced - github.com cloud default",
			args: protocol.RepoLookupArgs{
				Repo: api.RepoName("github.com/foo/bar"),
			},
			stored: []*types.Repo{},
			src:    repos.NewFakeSource(&githubSource, nil, githubRepository),
			result: &protocol.RepoLookupResult{Repo: &protocol.RepoInfo{
				ExternalRepo: api.ExternalRepoSpec{
					ID:          "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==",
					ServiceType: extsvc.TypeGitHub,
					ServiceID:   "https://github.com/",
				},
				Name:        "github.com/foo/bar",
				Description: "The description",
				VCS:         protocol.VCSInfo{URL: "git@github.com:foo/bar.git"},
				Links: &protocol.RepoLinks{
					Root:   "github.com/foo/bar",
					Tree:   "github.com/foo/bar/tree/{rev}/{path}",
					Blob:   "github.com/foo/bar/blob/{rev}/{path}",
					Commit: "github.com/foo/bar/commit/{commit}",
				},
			}},
			assert: typestest.AssertReposEqual(githubRepository),
		},
		{
			name: "found - github.com already exists",
			args: protocol.RepoLookupArgs{
				Repo: api.RepoName("github.com/foo/bar"),
			},
			stored: []*types.Repo{githubRepository},
			src:    repos.NewFakeSource(&githubSource, nil, githubRepository),
			result: &protocol.RepoLookupResult{Repo: &protocol.RepoInfo{
				ExternalRepo: api.ExternalRepoSpec{
					ID:          "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==",
					ServiceType: extsvc.TypeGitHub,
					ServiceID:   "https://github.com/",
				},
				Name:        "github.com/foo/bar",
				Description: "The description",
				VCS:         protocol.VCSInfo{URL: "git@github.com:foo/bar.git"},
				Links: &protocol.RepoLinks{
					Root:   "github.com/foo/bar",
					Tree:   "github.com/foo/bar/tree/{rev}/{path}",
					Blob:   "github.com/foo/bar/blob/{rev}/{path}",
					Commit: "github.com/foo/bar/commit/{commit}",
				},
			}},
		},
		{
			name: "not found - github.com",
			args: protocol.RepoLookupArgs{
				Repo: api.RepoName("github.com/foo/bar"),
			},
			src:    repos.NewFakeSource(&githubSource, github.ErrRepoNotFound),
			result: &protocol.RepoLookupResult{ErrorNotFound: true},
			err:    fmt.Sprintf("repository not found (name=%s notfound=%v)", api.RepoName("github.com/foo/bar"), true),
			assert: typestest.AssertReposEqual(),
		},
		{
			name: "unauthorized - github.com",
			args: protocol.RepoLookupArgs{
				Repo: api.RepoName("github.com/foo/bar"),
			},
			src:    repos.NewFakeSource(&githubSource, &github.APIError{Code: http.StatusUnauthorized}),
			result: &protocol.RepoLookupResult{ErrorUnauthorized: true},
			err:    fmt.Sprintf("not authorized (name=%s noauthz=%v)", api.RepoName("github.com/foo/bar"), true),
			assert: typestest.AssertReposEqual(),
		},
		{
			name: "temporarily unavailable - github.com",
			args: protocol.RepoLookupArgs{
				Repo: api.RepoName("github.com/foo/bar"),
			},
			src:    repos.NewFakeSource(&githubSource, &github.APIError{Message: "API rate limit exceeded"}),
			result: &protocol.RepoLookupResult{ErrorTemporarilyUnavailable: true},
			err: fmt.Sprintf(
				"repository temporarily unavailable (name=%s istemporary=%v)",
				api.RepoName("github.com/foo/bar"),
				true,
			),
			assert: typestest.AssertReposEqual(),
		},
		{
			name:   "synced - gitlab.com",
			args:   protocol.RepoLookupArgs{Repo: gitlabRepository.Name},
			stored: []*types.Repo{},
			src:    repos.NewFakeSource(&gitlabSource, nil, gitlabRepository),
			result: &protocol.RepoLookupResult{Repo: &protocol.RepoInfo{
				Name:        "gitlab.com/gitlab-org/gitaly",
				Description: "Gitaly is a Git RPC service for handling all the git calls made by GitLab",
				Fork:        false,
				Archived:    false,
				VCS: protocol.VCSInfo{
					URL: "https://gitlab.com/gitlab-org/gitaly.git",
				},
				Links: &protocol.RepoLinks{
					Root:   "https://gitlab.com/gitlab-org/gitaly",
					Tree:   "https://gitlab.com/gitlab-org/gitaly/tree/{rev}/{path}",
					Blob:   "https://gitlab.com/gitlab-org/gitaly/blob/{rev}/{path}",
					Commit: "https://gitlab.com/gitlab-org/gitaly/commit/{commit}",
				},
				ExternalRepo: gitlabRepository.ExternalRepo,
			}},
			assert: typestest.AssertReposEqual(gitlabRepository),
		},
		{
			name:   "found - gitlab.com",
			args:   protocol.RepoLookupArgs{Repo: gitlabRepository.Name},
			stored: []*types.Repo{gitlabRepository},
			src:    repos.NewFakeSource(&gitlabSource, nil, gitlabRepository),
			result: &protocol.RepoLookupResult{Repo: &protocol.RepoInfo{
				Name:        "gitlab.com/gitlab-org/gitaly",
				Description: "Gitaly is a Git RPC service for handling all the git calls made by GitLab",
				Fork:        false,
				Archived:    false,
				VCS: protocol.VCSInfo{
					URL: "https://gitlab.com/gitlab-org/gitaly.git",
				},
				Links: &protocol.RepoLinks{
					Root:   "https://gitlab.com/gitlab-org/gitaly",
					Tree:   "https://gitlab.com/gitlab-org/gitaly/tree/{rev}/{path}",
					Blob:   "https://gitlab.com/gitlab-org/gitaly/blob/{rev}/{path}",
					Commit: "https://gitlab.com/gitlab-org/gitaly/commit/{commit}",
				},
				ExternalRepo: gitlabRepository.ExternalRepo,
			}},
		},
		{
			name: "Private repos are not supported on sourcegraph.com",
			args: protocol.RepoLookupArgs{
				Repo: githubRepository.Name,
			},
			src: repos.NewFakeSource(&githubSource, nil, githubRepository.With(func(r *types.Repo) {
				r.Private = true
			})),
			result: &protocol.RepoLookupResult{ErrorNotFound: true},
			err:    fmt.Sprintf("repository not found (name=%s notfound=%v)", githubRepository.Name, true),
		},
		{
			name: "Private repos that used to be public should be removed asynchronously",
			args: protocol.RepoLookupArgs{
				Repo: githubRepository.Name,
			},
			src: repos.NewFakeSource(&githubSource, github.ErrRepoNotFound),
			stored: []*types.Repo{githubRepository.With(func(r *types.Repo) {
				r.UpdatedAt = r.UpdatedAt.Add(-time.Hour)
			})},
			result: &protocol.RepoLookupResult{Repo: &protocol.RepoInfo{
				ExternalRepo: api.ExternalRepoSpec{
					ID:          "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==",
					ServiceType: extsvc.TypeGitHub,
					ServiceID:   "https://github.com/",
				},
				Name:        "github.com/foo/bar",
				Description: "The description",
				VCS:         protocol.VCSInfo{URL: "git@github.com:foo/bar.git"},
				Links: &protocol.RepoLinks{
					Root:   "github.com/foo/bar",
					Tree:   "github.com/foo/bar/tree/{rev}/{path}",
					Blob:   "github.com/foo/bar/blob/{rev}/{path}",
					Commit: "github.com/foo/bar/commit/{commit}",
				},
			}},
			assertDelay: time.Second,
			assert:      typestest.AssertReposEqual(),
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			_, err := db.ExecContext(ctx, "DELETE FROM repo")
			if err != nil {
				t.Fatal(err)
			}

			rs := tc.stored.Clone()
			err = db.Repos().Create(ctx, rs...)
			if err != nil {
				t.Fatal(err)
			}

			clock := clock
			logger := logtest.Scoped(t)
			syncer := &syncer.Syncer{
				Now:     clock.Now,
				DB:      db,
				Sourcer: repos.NewFakeSourcer(nil, tc.src),
				ObsvCtx: observation.TestContextTB(t),
			}

			scheduler := scheduler.NewUpdateScheduler(logtest.Scoped(t), dbmocks.NewMockDB(), gitserver.NewMockClient())

			s := &Server{
				Logger:    logger,
				Syncer:    syncer,
				DB:        db,
				Scheduler: scheduler,
			}

			gs := grpc.NewServer(defaults.ServerOptions(logger)...)
			proto.RegisterRepoUpdaterServiceServer(gs, s)

			srv := httptest.NewServer(internalgrpc.MultiplexHandlers(gs, http.NotFoundHandler()))
			defer srv.Close()

			cli := repoupdater.NewClient(srv.URL)

			if tc.err == "" {
				tc.err = "<nil>"
			}

			res, err := cli.RepoLookup(ctx, tc.args)
			if have, want := fmt.Sprint(err), tc.err; have != want {
				t.Fatalf("have err: %q, want: %q", have, want)
			}

			if diff := cmp.Diff(res, tc.result, cmpopts.IgnoreFields(protocol.RepoInfo{}, "ID")); diff != "" {
				t.Fatalf("response mismatch(-have, +want): %s", diff)
			}

			if tc.assert != nil {
				if tc.assertDelay != 0 {
					time.Sleep(tc.assertDelay)
				}
				rs, err := db.Repos().List(ctx, database.ReposListOptions{})
				if err != nil {
					t.Fatal(err)
				}
				tc.assert(t, rs)
			}
		})
	}
}

type fakeScheduler struct{}

func (s *fakeScheduler) UpdateOnce(_ api.RepoID, _ api.RepoName) {}
func (s *fakeScheduler) ScheduleInfo(_ api.RepoID) *protocol.RepoUpdateSchedulerInfoResult {
	return &protocol.RepoUpdateSchedulerInfoResult{}
}
