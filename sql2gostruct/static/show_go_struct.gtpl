<!DOCTYPE html>
<html lang="en">
<head>
   <meta charset="UTF-8">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
   <title>Golang Code Viewer</title>
   <script src="/static/prism.min.js"></script>
   <script src="/static/copy.min.js"></script>
   <script src="/static/highlight.min.js"></script>
   <script src="/static/highlightjs-line-numbers.min.js"></script>
   <script>
    hljs.highlightAll(); // initialize highlighting
    hljs.initLineNumbersOnLoad(); // apply line numbering
   </script>
   <style>
       pre {
           background-color: #f0f0f0;
           padding: 10px;
           overflow-x: auto;
           white-space: pre-wrap;
           word-wrap: break-word;
       }
   </style>
</head>
<body>
   <pre><code class="language-go">{{.}}</code></pre>
   <button onclick="copyText()">复制代码</button>
   <script>
       function copyText() {
           const code = document.querySelector('.language-go').textContent;
           console.log(code);
           navigator.clipboard.writeText(code);
       }
   </script>
</body>
</html>