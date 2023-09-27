#!/usr/bin/env python3
import json
import os
import sys
import openai
from git import Repo
from unidiff import PatchSet


# 获取环境变量 OPENAI_API_KEY
def get_openai_key():
    key = os.getenv('OPENAI_API_KEY')
    if not key:
        print("not found OPENAI_API_KEY")
        sys.exit()
    return key


openai.api_key = get_openai_key()

sys_msg = '''
As a software development assistant, your role is to generate commit message for provided code diff.
The code diff will be in JSON format, comprising an array with three fields: file, added, and removed.
Please adhere to the following guidelines:
1. Commit messages should be limited to a maximum of 50 characters.
2. Prioritize actual code changes and disregard any irrelevant modifications.
'''


def git_diff():
    # 创建一个代表当前仓库的Repo对象
    repo = Repo()

    # 获取git diff的输出
    diffs = repo.git.diff()

    # 解析输出
    patch_set = PatchSet(diffs)

    files = []
    # 遍历每个差异
    for patch in patch_set:
        file = {'file': patch.path, 'added': '', 'removed': ''}
        for hunk in patch:
            for line in hunk:
                if line.value.strip() == "" or line.value.startswith('#'):
                    continue
                if line.is_added:
                    file['added'] += line.value
                elif line.is_removed:
                    file['removed'] += line.value
        files.append(file)
    return files


if __name__ == '__main__':
    diff_output = git_diff()
    print(json.dumps(diff_output))
    if len(diff_output) == 0:
        print("noting need to commit")
        sys.exit()

    if len(json.dumps(diff_output)) > 2**10:
        print("too long to waste token")
        sys.exit()

    completion = openai.ChatCompletion.create(
        model="gpt-3.5-turbo",
        messages=[
            {"role": "system", "content": sys_msg},
            {"role": "user", "content": json.dumps(diff_output)}
        ]
    )

    print(completion.choices[0].message)
    print(completion)
