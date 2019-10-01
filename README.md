# AWS reference architecture created by ansible and terraform
  
# Current state
- VPC role is completed
  
## VPC
- currently creates a VPC in the AZ defined in group_vars/all
- default AZ is us-east-1 with subnets created in all regions
  
### To Run
- Make a backup of secrets/vault.yml
    - ```cp secrets/vault.yml backup_vault.yml```
- Change the password in the vault file
- Decrypt and then update the group_vars/all file with your preferred values
    - ```ansible-vault rekey group_vars/all```
        - use the password from secrets/backup_vault.yml
        - when prompted for a new password, use the one you replaced in secrets/vault.yml
 


#### Default behavior is state=present
- Create the VPC and associated subnets:
	- ```ansible-playbook site.yml --vault-password-file secrets/vault.yml```
- Delete the VPC and subnets:
	- ```ansible-playbook site.yml -e state=absent --vault-password-file secrets/vault.yml```
  
## Internet Gateway
TODO
  
## Nat Gateways
TODO
  
# Bastion Host
TODO
