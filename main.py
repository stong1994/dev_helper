import json
import os
import sys
import openai

from gitxxx import git_diff


# 获取环境变量 OPENAI_API_KEY
def get_openai_key():
    key = os.getenv('OPENAI_API_KEY')
    if not key:
        print("not found OPENAI_API_KEY")
        sys.exit()
    return key


openai.api_key = get_openai_key()

sys_msg = '''
As a software developer assistant, your task is to write commit messages based on the given code.
Follow these guidelines:
1. A commit message should include a title and multiple body lines.
2. Adhere to best practices, such as keeping titles under 50 characters and limiting body lines to under 72 characters.
3. Utilize the diff output from user to create the summary
4. Diff output will use json string, it's an array and each item has 3 fields: file、added and removed.
'''



if __name__ == '__main__':
    diff_output = git_diff()
    print(json.dumps(diff_output))
    if len(diff_output) == 0:
        print("noting need to commit")
        sys.exit()

    completion = openai.ChatCompletion.create(
        model="gpt-3.5-turbo",
        messages=[
            {"role": "system", "content": "You are a helpful assistant."},
            {"role": "user", "content": json.dumps(diff_output)}
        ]
    )

    print(completion.choices[0].message)
    print(completion)