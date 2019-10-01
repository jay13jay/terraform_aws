# AWS reference architecture created by ansible and terraform
  
# Current state
- VPC role is completed
  
## VPC
- currently creates a VPC in the AZ defined in group_vars/all
- default AZ is us-east-1 with subnets created in all regions
  
### To Run
#### Default behavior is state=present
- Create the VPC and associated subnets:
	- ```ansible-playbook site.yml -e state=present```
- Delete the VPC and subnets:
	- ```ansible-playbook site.yml -e state=absent```
  
## Internet Gateway
TODO
  
## Nat Gateways
TODO
  
# Bastion Host
TODO