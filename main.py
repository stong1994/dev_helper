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
As a software development assistant, your role is to generate commit message for provided code diff. Please adhere to the following guidelines:

1. Commit message should consist of a concise title and ensure that the title is within 50 characters in length.
2. The code diff will be in JSON format, comprising an array with three fields: file, added, and removed.
3. Prioritize actual code changes and disregard any irrelevant modifications.
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
            {"role": "system", "content": sys_msg},
            {"role": "user", "content": json.dumps(diff_output)}
        ]
    )

    print(completion.choices[0].message)
    print(completion)
