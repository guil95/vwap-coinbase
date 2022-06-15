docker-build:
	 docker build -t vwap-coin .

docker-run:
	 docker run -i vwap-coin

docker-build-run:
	$(MAKE) docker-build
	$(MAKE) docker-run