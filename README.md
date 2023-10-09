# yaa
yaa - yaml search for humans 


yaa is designed to search the content of open source projects that use YAML files as their main file structure. For example, the SigmaHQ Project consists of approximately 2700 rules, each stored in a YAML file. Searching within the content of these files can be challenging. yaa provides an easy method, using query language, to search for specific content within these projects.

**Full Text Search**

search all rules that contain the keyword '7zip'. 

![image](https://github.com/alwashali/yaa/assets/22593441/85a9905c-6bb2-44e3-9e33-9b05f107322d)


## Query Language 
The power of yaa search comes from its query language, which depends on full-text search. Search can be as simple as writing a few keywords to search for any matches or to search inside a specific yaml property, or both. The default operation between search keywords is **OR.** 

 

**Basic search filters**

- +keywords means the word must appear (**AND** operation)
- -keywords means the keyword must not appears in the search result (**Not** operation)


![image](https://github.com/alwashali/yaa/assets/22593441/cb1ba680-b539-459d-92f7-b0f5e4317824)


## Examples

Search for **WMI** persistence related rules and exclude any rule having **WMI** in the falsepositive property. 

![image](https://github.com/alwashali/yaa/assets/22593441/8007a61b-7b91-483f-b330-b5ea45c336a8)



## Indexing yaml project 

yaa is built for searching inside detection rule projects such as SigmaHQ, however it can be used with any similar project. To index yaml project inside, use the command **index**, following example shows how to use the tool with SigmaHQ project. 


```
% git clone https://github.com/SigmaHQ/sigma.git
% ./yaa index sigma/rules/
```

![image](https://github.com/alwashali/yaa/assets/22593441/886d03f6-2120-4d22-a5e2-4530a68bf018)

## Export Matches 

```bash
% ./yaa search --export /tmp/sigma/ --limit 100  +title:hacktool 
79 files exported.
% ls /tmp/sigma | head
av_hacktool.yml
create_remote_thread_win_hktl_cactustorch.yml
create_remote_thread_win_hktl_cobaltstrike.yml
create_stream_hash_hacktool_download.yml
file_event_win_hktl_createminidump.yml
file_event_win_hktl_dumpert.yml
file_event_win_hktl_nppspy.yml
file_event_win_remote_access_tools_screenconnect_remote_file.yml
image_load_hktl_sharpevtmute.yml
image_load_hktl_silenttrinity_stager.yml

```


## Build yaa 

```bash
% git clone https://github.com/alwashali/yaa.git
% go build yaa.go
% ./yaa
NAME:
   Yaa - Yaml Searach for Humans

USAGE:
   Yaa [global options] command [command options] [arguments...]

COMMANDS:
   search, s  Search for sigma rules
   index, i   Path to yaml folder
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --limit value  Number of results to display (default: 10)
   --help, -h     show help
```





