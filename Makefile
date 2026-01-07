ANSIBLE_PLAYBOOK ?= ansible-playbook
ANSIBLE_ARGS ?= -K -i inventory/localhost.ini

.PHONY: bootstrap
bootstrap:
	@$(ANSIBLE_PLAYBOOK) bootstrap.yml $(ANSIBLE_ARGS)
