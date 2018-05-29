<html>
<head>
  <title>{{.StrConvert}}</title>
</head>
<body>
<h2>{{.StrConvert}}</h2>
<form action="/convertxml1to2file" method="post" enctype="multipart/form-data">
  <label for="file">Filename:</label>
  <input type="file" name="beerxml1file" id="beerxml1file">
  <br>
  <input type="submit" name="submit" value="Submit">
</form>
</body>
</html>