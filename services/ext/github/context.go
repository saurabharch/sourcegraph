package github

import (
	"strings"

	"github.com/sourcegraph/go-github/github"
	"golang.org/x/net/context"
	"sourcegraph.com/sourcegraph/sourcegraph/auth"
	"sourcegraph.com/sourcegraph/sourcegraph/server/serverctx"
	"sourcegraph.com/sourcegraph/sourcegraph/store"
	"sourcegraph.com/sourcegraph/sourcegraph/util/githubutil"
	"sourcegraph.com/sourcegraph/sourcegraph/util/traceutil"
)

func init() {
	// Make a GitHub API client available in the context that is
	// authenticated as the current user, or just using our
	// application credentials if there's no current user.
	//
	// This appends to LastFuncs, not just Funcs, because it must be
	// run AFTER the actor has been stored in the context, because it
	// depends on the actor.
	serverctx.LastFuncs = append(serverctx.LastFuncs,
		NewContextWithAuthedClient,
	)
}

type contextKey int

const (
	minimalClientKey contextKey = iota
)

// NewContextWithClient creates a new child context with the specified
// GitHub client.
func NewContextWithClient(ctx context.Context, client *github.Client, isAuthedUser bool) context.Context {
	return newContext(ctx, newMinimalClient(client, isAuthedUser))
}

// NewContextWithAuthedClient creates a new child context with a
// GitHub client that is authenticated using the credentials of the
// context's actor, or unauthenticated if there is no actor (or if the
// actor has no stored GitHub credentials).
func NewContextWithAuthedClient(ctx context.Context) (context.Context, error) {
	ghConf := *githubutil.Default
	ghConf.AppdashSpanID = traceutil.SpanIDFromContext(ctx)

	a := auth.ActorFromContext(ctx)
	var c *github.Client

	isAuthedUser := false
	if a.IsAuthenticated() {
		host := strings.TrimPrefix(githubutil.Default.BaseURL.Host, "api.") // api.github.com -> github.com
		tok, err := store.ExternalAuthTokensFromContext(ctx).GetUserToken(ctx, a.UID, host, githubutil.Default.OAuth.ClientID)
		if err == nil {
			c = ghConf.AuthedClient(tok.Token)
			isAuthedUser = true
		}
		if err != nil && err != auth.ErrNoExternalAuthToken && err != auth.ErrExternalAuthTokenDisabled {
			return nil, err
		}
	}
	if c == nil {
		c = ghConf.UnauthedClient()
	}
	return NewContextWithClient(ctx, c, isAuthedUser), nil
}

func newContext(ctx context.Context, client *minimalClient) context.Context {
	return context.WithValue(ctx, minimalClientKey, client)
}

// client returns the context's GitHub API client.
func client(ctx context.Context) *minimalClient {
	client, _ := ctx.Value(minimalClientKey).(*minimalClient)
	if client == nil {
		panic("no GitHub API client set in context")
	}
	return client
}
