import os


def readFns():
    ls=[]
    with open('./fns.txt') as f:
        line:str = f.readline()
        while line:
            if line != "":
                ls.append(line.strip().replace('\n',''))
            line = f.readline()
    return ls

f1=open('./dllimportdefs.go',"w")
f1.write("""
package dllimports

var dllImportDefs = []importTable{
""")


ls=readFns()
index=0
for l in ls:
    f1.write("""    /*{}*/ {{"{}", 0}},\n""".format(index,l))
    index=index+1
f1.write("}")

f1.write("\nconst (\n")
index=0
for l in ls:
    if l.startswith('_'):
        f1.write("""    {} = {}\n""".format(l.lstrip('_'), index))
    else:
        f1.write("""    {} = {}\n""".format(l, index))
    index=index+1
f1.write(")")




