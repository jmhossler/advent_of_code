.PHONY: check
check: lint test

.PHONY: test
test:
	pytest 2018/

.PHONY: lint
lint:
	pylint 2018/*.py --disable=C0111,C0103,E0401,R1710,C0200,R0903,R0912,R1702

get-%:
	./get_input.sh $*
