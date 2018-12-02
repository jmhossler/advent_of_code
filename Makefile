.PHONY: check
check: lint test

.PHONY: test
test:
	pytest 2018/

.PHONY: lint
lint:
	pylint 2018/*.py --disable=C0111,C0103,E0401,R1710,C0200

get-%:
	./get_input.sh $*
