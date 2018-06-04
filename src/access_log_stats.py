# -*- coding: utf8 -*-

import os

# 这个工具是用来统计tomcat和nginx的access, 计算url请求的qps, response的时间

#######################################
# file utils

LINE_BATCH = 1000

# 读取某个文件夹下面的文件夹数
def get_folders_byname(folder, name):
    fs = []
    path_dirs = os.listdir(folder)
    for path_dir in path_dirs:
        child = os.path.join("%s%s" % (folder, path_dir))
        if os.path.isdir(child):
            if child.find(name) != -1:
                fs.append(child)
    return fs

# 读取某个文件下面的文件
def get_files_byname(folder, name):
    files = []
    path_dirs = os.listdir(folder)
    for path_dir in path_dirs:
        child = os.path.join("%s/%s" % (folder, path_dir))
        if os.path.isfile(child):
            if child.find(name) != -1:
                files.append(child)
                print child
    return files

# 合并src文件到dest文件里面
def merge_file(src_files, dest_file):
    f_dest = open(dest_file, "w")
    lines = []
    for src_file in src_files:
        f_src = open(src_file, "r")
        line = f_src.readline()
        while line:
            lines.append(line)
            if len(lines) == LINE_BATCH:
                f_dest.writelines(lines)
                lines = []
            line = f_src.readline()
        if len(lines) > 0:
            f_dest.writelines(lines)
        f_src.close()
    f_dest.close()

#######################################
# nginx access files utils

def nginx_access_filter(access, result, filter):
    f_access = open(access, "r")
    f_result = open(result, "w")
    line = f_access.readline()
    while line:
        fields = line.split(" ")
        if fields[filter["pos"]] == filter["keyword"]:
            f_result.write(line)
        line = f_access.readline()
    f_access.close()
    f_result.close()

# 入口函数
if __name__ == "__main__":
    print "running..."
    # log_files = get_files_byname("C:/Joe/document/log_access", "access")
    # merge_file(log_files, "C:/Joe/document/log_access/merge_access")
    filter = {}
    filter["pos"] = 4
    filter["keyword"] = "status=500"
    nginx_access_filter("C:/Joe/document/log_access/merge_access","C:/Joe/document/log_access/filter_500_access", filter)
    filter["keyword"] = "status=200"
    nginx_access_filter("C:/Joe/document/log_access/merge_access", "C:/Joe/document/log_access/filter_200_access", filter)