ANSIBLE_PLAYBOOK ?= ansible-playbook
ANSIBLE_ARGS ?= -K

.PHONY: bootstrap
bootstrap:
	@$(ANSIBLE_PLAYBOOK) bootstrap.yml $(ANSIBLE_ARGS)
