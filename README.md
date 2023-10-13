# yaa
yaa - yaml search for humans 


yaa is a tool designed to search the content of open source projects that use YAML files as their primary file structure. It is primarily used for searching inside detection content projects but it can be used with any YAML-based project. SigmaHQ Project has approximately 2700 rules, and Splunk content library has approximately 1400 rules. Similarly, Nuclei templates and Sentinel detections, they contain a significant number of rules. Searching within the content of these files can be challenging, especially if you want a query language to extract specific search criteria. yaa provides a straightforward method by using a query language to search for specific content within the YAML files of these projects. 


## Query Language 
The power of yaa search comes from its query language, which depends on full-text search. Search can be as simple as writing a few keywords to search for any matches or to search inside a specific yaml property, or both. 

- +keywords means the word must appear (**AND** operation)
- -keywords means the keyword must not appears in the search result (**Not** operation)


### Full Text Search

Search all rules that contain the keyword '7zip'. 

![image](https://github.com/alwashali/yaa/assets/22593441/85a9905c-6bb2-44e3-9e33-9b05f107322d)

 
Yaml property search with AND operator, searching for any rule that has the word 'powershell **AND** obfuscation' in the title property.

![image](https://github.com/alwashali/yaa/assets/22593441/cb1ba680-b539-459d-92f7-b0f5e4317824)


### Exclude Result

Search for **WMI** persistence related rules and exclude any rule having **WMI** in the falsepositive property. 

![image](https://github.com/alwashali/yaa/assets/22593441/8007a61b-7b91-483f-b330-b5ea45c336a8)


### Nested property search

Yaa can index nested properties and make them searchable by specifying the nested property name prefixed with a dot.

![image](https://github.com/alwashali/yaa/assets/22593441/b5ea4e28-b481-4277-b308-7d0b536b1d69)



### Export Matches

Files matching the search criteria can be exported to a differnet directory

 ![image](https://github.com/alwashali/yaa/assets/22593441/ca5f6433-0b24-4ad7-b495-26bd67ff8354)




### Indexing yaml project 

yaa is built for searching inside detection rule projects such as SigmaHQ, however it can be used with any similar project. To index a yaml project, use the command **index**. 


```
% git clone https://github.com/SigmaHQ/sigma.git
% ./yaa index sigma/rules/
```

![image](https://github.com/alwashali/yaa/assets/22593441/886d03f6-2120-4d22-a5e2-4530a68bf018)



### Build yaa 

```bash
% git clone https://github.com/alwashali/yaa.git
% go build yaa.go
% ./yaa

NAME:
   Yaa - Yaml Searach for Humans

USAGE:
   Yaa [global options] command [command options] [arguments...]

COMMANDS:
   search, s  
   index, i   Path to yaml folder
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```





