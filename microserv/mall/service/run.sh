# /bin/bash
cd user/rpc
go run user.go &
cd ../api
go run user.go &
cd ../../product/rpc
go run product.go &
cd ../api
go run product.go &
cd ../../order/rpc
go run order.go &
cd ../api
go run order.go &