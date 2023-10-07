	"sort"
	"google.golang.org/protobuf/encoding/protojson"
				return &gitserver.MockGRPCClient{
					MockRepoDelete: mockRepoDelete,
					return &gitserver.MockGRPCClient{MockP4Exec: mockP4Exec}
			return &gitserver.MockGRPCClient{MockBatchLog: mockBatchLog}
			return &gitserver.MockGRPCClient{
				MockBatchLog: mockBatchLog,
					return &gitserver.MockGRPCClient{MockIsRepoCloneable: mockIsRepoCloneable}
					return &gitserver.MockGRPCClient{MockIsRepoCloneable: mockIsRepoCloneable}
	const (
		gitserverAddr1 = "172.16.8.1:8080"
		gitserverAddr2 = "172.16.8.2:8080"
	)

	expectedResponses := []gitserver.SystemInfo{
		{
			Address:    gitserverAddr1,
			FreeSpace:  102400,
			TotalSpace: 409600,
		},
		{
			Address:    gitserverAddr2,
			FreeSpace:  51200,
			TotalSpace: 204800,
		},
	sort.Slice(expectedResponses, func(i, j int) bool {
		return expectedResponses[i].Address < expectedResponses[j].Address
	})

		sort.Slice(info, func(i, j int) bool {
			return info[i].Address < info[j].Address
		})

		require.Len(t, info, len(expectedResponses), "expected %d disk info(s)", len(expectedResponses))
		for i := range expectedResponses {
			require.Equal(t, info[i].Address, expectedResponses[i].Address)
			require.Equal(t, info[i].FreeSpace, expectedResponses[i].FreeSpace)
			require.Equal(t, info[i].TotalSpace, expectedResponses[i].TotalSpace)
		}
		source := gitserver.NewTestClientSource(t, []string{gitserverAddr1, gitserverAddr2}, func(o *gitserver.TestClientSourceOptions) {
			responseByAddress := make(map[string]*proto.DiskInfoResponse, len(expectedResponses))
			for _, response := range expectedResponses {
				responseByAddress[response.Address] = &proto.DiskInfoResponse{
					FreeSpace:   response.FreeSpace,
					TotalSpace:  response.TotalSpace,
					PercentUsed: response.PercentUsed,
				}
			}

					address := cc.Target()
					response, ok := responseByAddress[address]
					if !ok {
						t.Fatalf("received unexpected address %q", address)
					}

					return response, nil
				return &gitserver.MockGRPCClient{MockDiskInfo: mockDiskInfo}
		source := gitserver.NewTestClientSource(t, []string{gitserverAddr1, gitserverAddr2}, func(o *gitserver.TestClientSourceOptions) {
					return nil, nil
				return &gitserver.MockGRPCClient{MockDiskInfo: mockDiskInfo}
				responseByAddress := make(map[string]*proto.DiskInfoResponse, len(expectedResponses))
				for _, response := range expectedResponses {
					responseByAddress[fmt.Sprintf("http://%s/disk-info", response.Address)] = &proto.DiskInfoResponse{
						FreeSpace:   response.FreeSpace,
						TotalSpace:  response.TotalSpace,
						PercentUsed: response.PercentUsed,
					}
				}

				address := r.URL.String()
				response, ok := responseByAddress[address]
				if !ok {
					return nil, errors.Newf("unexpected URL: %q", address)

				encoded, _ := protojson.Marshal(response)
				body := io.NopCloser(strings.NewReader(string(encoded)))
				return &http.Response{
					StatusCode: 200,
					Body:       body,
				}, nil
				return &gitserver.MockGRPCClient{MockDiskInfo: mockDiskInfo}
				return &gitserver.MockGRPCClient{MockDiskInfo: mockDiskInfo}