# for devlopment
```bash
cd src/
docker build -t xyz .
docker run -it --rm -p 80:8080 -v ${PWD}:/app xyz bash
go build -o main && ./main
```

# for production
```bash
cd src/
docker build -t xyz .
docker run --rm -p 80:8080 xyz
```

# References
[Charts](https://blog.logrocket.com/visualizing-data-go-echarts/)