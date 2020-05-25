```bash
# command-line-arguments
./example.go:37:29: cannot convert conn (type *grpc.ClientConn) to type cacheService.CacheServiceClient:
	*grpc.ClientConn does not implement cacheService.CacheServiceClient (missing Add method)
./example.go:87:59: invalid indirect of emptypb.Empty literal (type emptypb.Empty)
```