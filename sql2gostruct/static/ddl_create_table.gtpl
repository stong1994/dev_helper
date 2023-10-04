
<html>
<head>
<title></title>
<script src="/static/jquery-3.6.0.min.js"></script>
</head>
<body>

<form action="/" method="post" data-dismiss="true">
 <p><label for="ddl">输入ddl:</label></p>
 <textarea id="ddl" name="ddl" rows="16" cols="50"></textarea>
 <br>
 <input type="submit" value="提交">
</form>

<script>
    $('#ddl').on('submit', function(e) {
      e.preventDefault(); // 阻止表单的提交操作
      $('#ddl').submit(); // 使用 .submit() 方法代替 .submit() 方法
  });
</script>
</body>
</html>
