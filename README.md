# yaa
yaa - yaml search for humans 


yaa is designed to search the content of open source projects that use YAML files as their main file structure. For example, the SigmaHQ Project consists of approximately 2700 rules, each stored in a YAML file. Searching within the content of these files can be challenging. yaa provides an easy method, using query language, to search for specific content within these projects.

**Full Text Search**

![image](https://github.com/alwashali/yaa/assets/22593441/85a9905c-6bb2-44e3-9e33-9b05f107322d)

**Query Language** 
The power of yaa search comes from its query language, which depends on full-text search. Search can be as simple as writing a few keywords to search for any matches or to search inside a specific yaml property, or both. The default operation between search keywords is **OR.** 

 

**Basic search filters**

- +keywords means the word must appear (**AND** operation)
- -keywords means the keyword must not appears in the search result (**Not** operation)


![image](https://github.com/alwashali/yaa/assets/22593441/cb1ba680-b539-459d-92f7-b0f5e4317824)



Using NOT operation to exclude result having **WMI** in falsepositive property, and limiting the display result to 3 rules only. 

![image](https://github.com/alwashali/yaa/assets/22593441/8007a61b-7b91-483f-b330-b5ea45c336a8)



## Indexing yaml project 

yaa is built for searching inside detection rule projects such as SigmaHQ, howver it can be used with any similar project. 


```
% git clone https://github.com/SigmaHQ/sigma.git
% ./yaa index sigma/rules/
```

![image](https://github.com/alwashali/yaa/assets/22593441/886d03f6-2120-4d22-a5e2-4530a68bf018)



