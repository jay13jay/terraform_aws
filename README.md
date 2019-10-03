# AWS reference architecture created by ansible and terraform
  
# Current state
- VPC role is completed

# Build State
[![Build Status](https://codebuild.us-east-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoiUUMxeGJMVkd6Nmx1TDBhTC91RlAramd2anZsc1YrYVlod3N4MEtUTDdYMHNCd3hON0RXSFl5RzMrT044S3NTM09YT3VLZmhPRkZ1eWNXWkdUNC9QZllvPSIsIml2UGFyYW1ldGVyU3BlYyI6IlFtbVloSzN3bmh3ZTN1T1AiLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)](https://github.com/jay13jay/terraform_aws)
  
## VPC
- currently creates a VPC in the AZ defined in group_vars/all
- default AZ is us-east-1 with subnets created in all regions
    - A public and Private /24 subnet will be created in each region
    - provides 251 hosts maximum. a different subnet scheme will be needed for more nodes per region
  
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
