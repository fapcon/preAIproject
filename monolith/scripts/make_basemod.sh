echo 'FROM golang:1.19.2-alpine' > Dockerfile
echo 'ADD ./download /go/pkg/mod/cache/download' >> Dockerfile
mkdir download
cp -r $GOPATH/pkg/mod/cache/download/studentgit.kata.academy ./download
docker build -t eazzygroup/basemod .
cat dockerpass.txt | docker login --username studentkata --password-stdin
docker push eazzygroup/basemod
