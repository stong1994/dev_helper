import os
import openai
from dotenv import load_dotenv, find_dotenv
import subprocess


# 读取本地/项目的环境变量。

# find_dotenv()寻找并定位.env文件的路径
# load_dotenv()读取该.env文件，并将其中的环境变量加载到当前的运行环境中
# 如果你设置的是全局的环境变量，这行代码则没有任何作用。
_ = load_dotenv(find_dotenv())


# 获取环境变量 OPENAI_API_KEY
def get_openai_key():
    _ = load_dotenv(find_dotenv())
    return os.environ['OPENAI_API_KEY']


openai.api_key = get_openai_key()

system_msg = '''
As a software developer assistant, your task is to write commit messages based on the given code.
Follow these guidelines:
1. A commit message should include a title and multiple body lines.
2. Adhere to best practices, such as keeping titles under 50 characters and limiting body lines to under 72 characters.

'''

def get_git_diff():
    # 运行git diff命令并获取输出
    result = subprocess.run(['git', 'diff'], stdout=subprocess.PIPE)
    # 将输出从字节转换为字符串
    diff_output = result.stdout.decode('utf-8')
    return diff_output

def parse_git_diff(diff_output):
    # 将diff输出分割成行
    lines = diff_output.split('\n')
    # 分别处理输入和输出
    for line in lines:
        if line.startswith('-'):
            print('Input: {line[1:]}')
        elif line.startswith('+'):
            print('Output: {line[1:]}')

diff_output = get_git_diff()
parse_git_diff(diff_output)


if __name__ == '__main__':
    openai.ChatCompletion.create(
        model="gpt-3.5-turbo",
        messages=[
            {"role": "system", "content": system_msg},
            {"role": "user", "content": "Who won the world series in 2020?"},
            {"role": "assistant", "content": "The Los Angeles Dodgers won the World Series in 2020."},
            {"role": "user", "content": "Where was it played?"}
        ]
    )