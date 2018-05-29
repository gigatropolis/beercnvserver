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

<h>File converted to XML 2.0</h>
<div>Download XML2 file: <a href="/downloads/xml2/{{.FileName}}" download>{{.FileName}}</a></div>
<div>
	<pre>{{.XML2}}</pre>
</div>
</body>
</html>