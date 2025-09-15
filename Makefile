ANSIBLE_PLAYBOOK ?= ansible-playbook
ANSIBLE_INVENTORY ?= inventories/localhost.yml
ANSIBLE_ARGS ?= -K

.PHONY: bootstrap
bootstrap:
	@$(ANSIBLE_PLAYBOOK) -i $(ANSIBLE_INVENTORY) bootstrap.yml $(ANSIBLE_ARGS)
