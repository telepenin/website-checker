.EXPORT_ALL_VARIABLES:

# Current version
VERSION ?= 0.0.1

# repo for images
IMG_REPO = quay.io/telepenin


run-consumer:
	cd consumer/; $(MAKE) run; cd -;

run-producer:
	cd producer/; $(MAKE) run; cd -;