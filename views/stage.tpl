<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta name="viewport" content="width=device-width,minimum-scale=1.0,maximum-scale=1.0,user-scalable=no">
<title>道具管理后台</title>
<style>
*{margin: 0;padding: 0;}
			body {
				padding: 40px 100px;
			}
			.demo {
			width: 600px;
			margin: 40px auto;
			font-family: 'trebuchet MS', 'Lucida sans', Arial;
			font-size: 14px;
			color: #444;
			}

			table {
				*border-collapse: collapse; /* IE7 and lower */
				border-spacing: 0;
				width: 100%;
				border-spacing: 0;
			}
			/*========bordered table========*/
			.bordered {
				border: solid #ccc 1px;
				-moz-border-radius: 6px;
				-webkit-border-radius: 6px;
				border-radius: 6px;
				-webkit-box-shadow: 0 1px 1px #ccc;
				-moz-box-shadow: 0 1px 1px #ccc;
				box-shadow: 0 1px 1px #ccc;
			}

			.bordered tr {
				-o-transition: all 0.1s ease-in-out;
				-webkit-transition: all 0.1s ease-in-out;
				-moz-transition: all 0.1s ease-in-out;
				-ms-transition: all 0.1s ease-in-out;
				transition: all 0.1s ease-in-out;		
			}
			.bordered .highlight,
			.bordered tr:hover {
				background: #fbf8e9;		
			}
			.bordered td, 
			.bordered th {
				border-left: 1px solid #ccc;
				border-top: 1px solid #ccc;
				padding: 10px;
				text-align: left;
			}
			.bordered th {
				background-color: #dce9f9;
				background-image: -webkit-gradient(linear, left top, left bottom, from(#ebf3fc), to(#dce9f9));
				background-image: -webkit-linear-gradient(top, #ebf3fc, #dce9f9);
				background-image: -moz-linear-gradient(top, #ebf3fc, #dce9f9);
				background-image: -ms-linear-gradient(top, #ebf3fc, #dce9f9);
				background-image: -o-linear-gradient(top, #ebf3fc, #dce9f9);
				background-image: linear-gradient(top, #ebf3fc, #dce9f9);
				filter: progid:DXImageTransform.Microsoft.gradient(GradientType=0, startColorstr=#ebf3fc, endColorstr=#dce9f9);
				-ms-filter: "progid:DXImageTransform.Microsoft.gradient (GradientType=0, startColorstr=#ebf3fc, endColorstr=#dce9f9)";
				-webkit-box-shadow: 0 1px 0 rgba(255,255,255,.8) inset;
				-moz-box-shadow:0 1px 0 rgba(255,255,255,.8) inset;
				box-shadow: 0 1px 0 rgba(255,255,255,.8) inset;
				border-top: none;
				text-shadow: 0 1px 0 rgba(255,255,255,.5);
			}
			.bordered td:first-child, 
			.bordered th:first-child {
				border-left: none;
			}
			.bordered th:first-child {
				-moz-border-radius: 6px 0 0 0;
				-webkit-border-radius: 6px 0 0 0;
				border-radius: 6px 0 0 0;
			}
			.bordered th:last-child {
				-moz-border-radius: 0 6px 0 0;
				-webkit-border-radius: 0 6px 0 0;
				border-radius: 0 6px 0 0;
			}
			.bordered tr:last-child td:first-child {
				-moz-border-radius: 0 0 0 6px;
				-webkit-border-radius: 0 0 0 6px;
				border-radius: 0 0 0 6px;
			}
			.bordered tr:last-child td:last-child {
				-moz-border-radius: 0 0 6px 0;
				-webkit-border-radius: 0 0 6px 0;
				border-radius: 0 0 6px 0;
			}
			/*----------------------*/
			.zebra td, 
			.zebra th {
				padding: 10px;
				border-bottom: 1px solid #f2f2f2;
			}
			.zebra .alternate,
			.zebra tbody tr:nth-child(even) {
				background: #f5f5f5;
				-webkit-box-shadow: 0 1px 0 rgba(255,255,255,.8) inset;
				-moz-box-shadow:0 1px 0 rgba(255,255,255,.8) inset;
				box-shadow: 0 1px 0 rgba(255,255,255,.8) inset;
			}
			.zebra th {
				text-align: left;
				text-shadow: 0 1px 0 rgba(255,255,255,.5);
				border-bottom: 1px solid #ccc;
				background-color: #eee;
				background-image: -webkit-gradient(linear, left top, left bottom, from(#f5f5f5), to(#eee));
				background-image: -webkit-linear-gradient(top, #f5f5f5, #eee);
				background-image: -moz-linear-gradient(top, #f5f5f5, #eee);
				background-image: -ms-linear-gradient(top, #f5f5f5, #eee);
				background-image: -o-linear-gradient(top, #f5f5f5, #eee);
				background-image: linear-gradient(top, #f5f5f5, #eee);
				filter: progid:DXImageTransform.Microsoft.gradient(GradientType=0, startColorstr=#f5f5f5, endColorstr=#eeeeee);
				-ms-filter: "progid:DXImageTransform.Microsoft.gradient (GradientType=0, startColorstr=#f5f5f5, endColorstr=#eeeeee)";
			}
			.zebra th:first-child {
				-moz-border-radius: 6px 0 0 0;
				-webkit-border-radius: 6px 0 0 0;
				border-radius: 6px 0 0 0;
			}
			.zebra th:last-child {
				-moz-border-radius: 0 6px 0 0;
				-webkit-border-radius: 0 6px 0 0;
				border-radius: 0 6px 0 0;
			}
			.zebra tfoot td {
				border-bottom: 0;
				border-top: 1px solid #fff;
				background-color: #f1f1f1;
			}
			.zebra tfoot td:first-child {
				-moz-border-radius: 0 0 0 6px;
				-webkit-border-radius: 0 0 0 6px;
				border-radius: 0 0 0 6px;
			}
			.zebra tfoot td:last-child {
				-moz-border-radius: 0 0 6px 0;
				-webkit-border-radius: 0 0 6px 0;
				border-radius: 0 0 6px 0;
			}


			.btn{
		display:inline-block;
		zoom:1;
		*display:inline;
		vertical-align:baseline;
		margin:0 2px;
		outline:none;
		overflow:visible;
		width:auto;
		*width:1;
		cursor:pointer;
		text-align:center;
		text-decoration:none;
		font:14px/100% Arial, Helvetica, sans-serif;
		padding:0.5em 2em .55em;
		text-shadow:0 1px 1px rgba(0,0,0.3);
		-webkit-border-radius:.5em;
		-moz-border-radius:.5em;
		border-radius:.5em;
		border:0 none;
		-webkit-box-shadow:0 1px 2px rgba(0,0,0,.2);
		-moz-box-shadow:0 1px 2px rgba(0,0,0,.2);
		box-shadow:0 1px 2px rgba(0,0,0,.2);
		-webkit-transition:all 0.5s ease 0.218s;
		-moz-transition: all 0.5s ease 0.218s;
		transition: all 0.5s ease 0.218s;
	}
.btn:hover{
		text-decoration:none;
	}	
.btn:active{
		position:relative;
		top:1px;
	}



	.primary {
					color: #fff;
					border: solid 1px #da7c0c;
					background: #f47a20;
					background-repeat: repeat-x;
					background-image: -khtml-gradient(linear, left top, left bottom, from(#faa51a), to(#f47a20));
					background-image: -moz-linear-gradient(#faa51a, #f47a20);
					background-image: -ms-linear-gradient(#faa51a, #f47a20);
					background-image: -webkit-gradient(linear, left top, left bottom, color-stop(0%, #faa51a), color-stop(100%, #f47a20));
					background-image: -webkit-linear-gradient(#faa51a, #f47a20);
					background-image: -o-linear-gradient(#faa51a, #f47a20);
					background-image: linear-gradient(#faa51a, #f47a20);
					filter: progid:DXImageTransform.Microsoft.gradient(GradientType=0, startColorstr=#faa51a, endColorstr=#f47a20);/*IE<9>*/
					-ms-filter: "progid:DXImageTransform.Microsoft.gradient (GradientType=0, startColorstr=#faa51a, endColorstr=#f47a20)";/*IE8+*/
				}
				.primary:hover {
					background: #f06015;
					background-repeat: repeat-x;
					background-image: -khtml-gradient(linear, left top, left bottom, from(#f88e11), to(#f06015));
					background-image: -moz-linear-gradient(#f88e11, #f06015);
					background-image: -ms-linear-gradient(#f88e11, #f06015);
					background-image: -webkit-gradient(linear, left top, left bottom, color-stop(0%, #f88e11), color-stop(100%, #f06015));
					background-image: -webkit-linear-gradient(#f88e11, #f06015);
					background-image: -o-linear-gradient(#f88e11, #f06015);
					background-image: linear-gradient(#f88e11, #f06015);
					filter: progid:DXImageTransform.Microsoft.gradient(GradientType=0, startColorstr=#f88e11, endColorstr=#f06015);/*IE<9>*/
					-ms-filter: "progid:DXImageTransform.Microsoft.gradient (GradientType=0, startColorstr=#f88e11, endColorstr=#f06015)";/*IE8+*/
				}
				
				.primary:active {
					color: #fcd3a5;
					background: #faa51a;
					background-repeat: repeat-x;
					background-image: -khtml-gradient(linear, left top, left bottom, from(#f47a20), to(#faa51a));
					background-image: -moz-linear-gradient(#f47a20, #faa51a);
					background-image: -ms-linear-gradient(#f47a20, #faa51a);
					background-image: -webkit-gradient(linear, left top, left bottom, color-stop(0%, #f47a20), color-stop(100%, #faa51a));
					background-image: -webkit-linear-gradient(#f47a20, #faa51a);
					background-image: -o-linear-gradient(#f47a20, #faa51a);
					background-image: linear-gradient(#f47a20, #faa51a);
					filter: progid:DXImageTransform.Microsoft.gradient(GradientType=0, startColorstr=#f47a20, endColorstr=#faa51a);/*IE<9>*/
					-ms-filter: "progid:DXImageTransform.Microsoft.gradient (GradientType=0, startColorstr=#f47a20, endColorstr=#faa51a)";/*IE8+*/
				}

	/*输入框*/
	/*.search{
	position:relative;		
	}*/
/*.search:before{
	content:"";
	border:1px solid #777;
	border-width:1px 1px 2px;
	width:5px;
	height:0;
	display:inline-block;
	-moz-transform:rotate(45deg);
	-webkit-transform:rotate(45deg);
	-o-transform:rotate(45deg);
	-ms-transform:rotate(45deg);
	transform:rotate(45deg);
	position:absolute;
	left:16px;
	top:15px;
	z-index:3;	
	}	*/
.search:after {
			content: "";
			width: 5px;
			height: 5px;
			border: 2px solid #777;
			-webkit-border-radius: 5px;
			-moz-border-radius: 5px;
			border-radius: 5px;
			position: absolute;
			z-index: 2;
			left: 10px;
			top: 7px;
			display: inline-block;
		}
		
		.search input[type="text"] {
			font: bold 12px Arial,Helvetica,Sans-serif;
			width: 150px;
			padding: 6px 15px 6px 35px;
			border: 0 none;
			color: #777;
			-moz-border-radius: 20px;
			-webkit-border-radius: 20px;
			border-radius: 20px;
			-webkit-transition: all 0.7s ease 0s;
			-moz-transition: all 0.7s ease 0s;
			-o-transition: all 0.7s ease 0s;
			transition: all 0.7s ease 0s;
		}
		.search input[type="number"] {
			font: bold 12px Arial,Helvetica,Sans-serif;
			width: 150px;
			padding: 6px 15px 6px 35px;
			border: 0 none;
			color: #777;
			-moz-border-radius: 20px;
			-webkit-border-radius: 20px;
			border-radius: 20px;
			-webkit-transition: all 0.7s ease 0s;
			-moz-transition: all 0.7s ease 0s;
			-o-transition: all 0.7s ease 0s;
			transition: all 0.7s ease 0s;
		}
		/*.search input[type="text"]:focus {
			width: 200px;
		}*/
/*===========黑色背景色的Search Box============*/
		.bgBlack input[type="text"] {
			background-color: #d0d0d0;
			text-shadow: 0 2px 2px rgba(0, 0, 0, 0.3);
			-webkit-box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 3px rgba(0, 0, 0, 0.2) inset;
			-moz-box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 3px rgba(0, 0, 0, 0.2) inset;
			box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 3px rgba(0, 0, 0, 0.2) inset;    
		}
/*===============高亮背景色的Search Box==============*/
		.bgLight input[type="text"]{
			background-color: #fcfcfc;/*change the background color*/
			color: #bebebe;/*change the font color*/
			text-shadow: 0 2px 3px rgba(0, 0, 0, 0.1);
			-webkit-box-shadow: 0 1px 3px rgba(0, 0, 0, 0.15) inset;
			-moz-box-shadow: 0 1px 3px rgba(0, 0, 0, 0.15) inset;
			box-shadow: 0 1px 3px rgba(0, 0, 0, 0.15) inset;
		}
		.bgLight:after,
		.bgLight:before {
			border-color: #bebebe;/*change the icon color*/
		}
	/*==============Apple.com Search Box效果====================*/
		.appleSearch input[type="text"] {
			background-color: #444;
			color: #d7d7d7;
			text-shadow: 0 2px 2px rgba(0, 0, 0, 0.3);
			-webkit-box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 3px rgba(0, 0, 0, 0.2) inset;
			-moz-box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 3px rgba(0, 0, 0, 0.2) inset;
			box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 3px rgba(0, 0, 0, 0.2) inset;    
		}
		
		.appleSearch input[type="text"]:focus {
			background-color: #fcfcfc;
			color: #6a6f75;
			-webkit-box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 0 rgba(0, 0, 0, 0.9) inset;
			-moz-box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 0 rgba(0, 0, 0, 0.9) inset;
			box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 0 rgba(0, 0, 0, 0.9) inset;
			text-shadow: 0 2px 3px rgba(0, 0, 0, 0.1);
    }

    /*===========黑色背景色的Search Box============*/
		.bgBlack input[type="number"] {
			background-color: #d0d0d0;
			text-shadow: 0 2px 2px rgba(0, 0, 0, 0.3);
			-webkit-box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 3px rgba(0, 0, 0, 0.2) inset;
			-moz-box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 3px rgba(0, 0, 0, 0.2) inset;
			box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 3px rgba(0, 0, 0, 0.2) inset;    
		}
/*===============高亮背景色的Search Box==============*/
		.bgLight input[type="number"]{
			background-color: #fcfcfc;/*change the background color*/
			color: #bebebe;/*change the font color*/
			text-shadow: 0 2px 3px rgba(0, 0, 0, 0.1);
			-webkit-box-shadow: 0 1px 3px rgba(0, 0, 0, 0.15) inset;
			-moz-box-shadow: 0 1px 3px rgba(0, 0, 0, 0.15) inset;
			box-shadow: 0 1px 3px rgba(0, 0, 0, 0.15) inset;
		}
		.bgLight:after,
		.bgLight:before {
			border-color: #bebebe;/*change the icon color*/
		}
	/*==============Apple.com Search Box效果====================*/
		.appleSearch input[type="number"] {
			background-color: #444;
			color: #d7d7d7;
			text-shadow: 0 2px 2px rgba(0, 0, 0, 0.3);
			-webkit-box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 3px rgba(0, 0, 0, 0.2) inset;
			-moz-box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 3px rgba(0, 0, 0, 0.2) inset;
			box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 3px rgba(0, 0, 0, 0.2) inset;    
		}
		
		.appleSearch input[type="number"]:focus {
			background-color: #fcfcfc;
			color: #6a6f75;
			-webkit-box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 0 rgba(0, 0, 0, 0.9) inset;
			-moz-box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 0 rgba(0, 0, 0, 0.9) inset;
			box-shadow: 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 0 rgba(0, 0, 0, 0.9) inset;
			text-shadow: 0 2px 3px rgba(0, 0, 0, 0.1);
    }
#form{
		background:#666;
	}			

a{
	text-decoration: NONE
}
</style>
</head>
<body>
	<h1>奇妙的朋友管理后台</h1>
	<hr/>
<center>
<!--添加道具-->
<br />
<h2>道具管理</h2>
<form action="/stage" method="POST" onsubmit="return checkselect();" class="search bgBlack">
	<table class="bordered">
	<tr><td>游戏账号：</td><td><input type="text"  name="useracount" id="useracount" placeholder="玩家的id号"/></td></tr>
	<tr><td>赠送道具的类型：</td><td>
	<select id="mySelect" name="stagename" onchange="cc(this[selectedIndex].value);">
		<option value="0">请选择附加类型</option>
		<option value="1">钻石</option>
		<option value="2">金币</option>
		<option value="3">行动力（花）</option>
		<option value="4">道具</option>

	</select></td></tr>
	<tr><td>赠送道具的子类型：</td><td>
	<select  id="mySelect_sub" name="stagename_sub">
		<option value="0">请选择附加类型</option>
		<option value="1">快速消除</option>
		<option value="2">重新排列</option>
		<option value="3">交换位置</option>
		<option value="4">旋风射线</option>
		<option value="5">加五步</option>

	</select></td></tr>
	<tr><td>赠送道具的数量：</td><td><input type="number" min=0  name="stagenum" id="stagenum" placeholder="道具的数量"/></td></tr><br />
	<tr><td  colspan="2"> <input type="submit" name="提交" value="提交" class="btn primary"/>
	<input type="reset" name="取消" value="取消" class="btn primary"/></td></tr>
	
</table>
</form>
<!--编辑消息公告-->
<br />
<h2>公告编辑板块  <a href="/message">管理</a></h2>

<form action="/message" method="POST" class="search bgBlack">
	<table class="bordered">
		<tr><td>标题：</td><td><input type="text" name="title"/></td></tr>
		<tr><td>公告内容：</td><td><textarea rows="10" cols="40" style="resize:none" name="content"></textarea></td></tr>
		<tr><td>标题2：</td><td><input type="text" name="title2"/></td></tr>
		<tr><td>公告内容2：</td><td><textarea rows="10" cols="40" style="resize:none" name="content2"></textarea></td></tr>
		<tr><td  colspan="2"><input type="submit" name="提交" value="提交" class="btn primary"/>
		<input type="reset" name="取消" value="取消" class="btn primary"/></td></tr>
	</table>
</form>
<br/>
<br/>
<table class="bordered">
	<tr><th>用户账号</th><th>道具</th><th>道具数量</th><th>操作时间</th></tr>
{{range $k, $v := .s}}


<tr><td>{{$v.Useracount}}</td><td>
	{{if eq $v.Stagename "1"}}
		钻石
	
	{{else if eq $v.Stagename "2"}}
		金币
	
	{{else if eq $v.Stagename "3"}}
		行动力（花）
	
	{{else if eq $v.Stagename "4"}}
		{{if eq $v.Stagenamesub "1"}}
			快速消除
		
		{{else if eq $v.Stagenamesub "2"}}
			重新排列
		
		{{else if eq $v.Stagenamesub "3"}}
			交换位置
		
		{{else if eq $v.Stagenamesub "4"}}
			旋风射线
		{{else if eq $v.Stagenamesub "5"}}
			加五步
		{{else}}
			选择无效
		{{end}}
	
	{{else}}
		选择无效
	{{end}}
</td><td>{{$v.Stagenum}}</td><td>{{$v.Time}}</td></tr>

{{end}} 
</table>
</center>
<script>
	
	function checkselect(){
	 	var t = document.getElementById("mySelect");
	 	
	 	var  useracount=document.getElementById("useracount").value;
	 	var stagenum=document.getElementById("stagenum").value;
		value=t.options[t.selectedIndex].value
		if (value=="0"||useracount==""||stagenum==""){

			return false;

		}else if (value=="4"){
			val=checkselect_sub();
			if (val=="0"){
				return false;
			}else{
				return true;
			}
		}else{
			return true
		}
	}
	function checkselect_sub(){
		var t=document.getElementById("mySelect_sub");
		value=t.options[t.selectedIndex].value
		return value
	}
	function cc(val){
		var mySelect_sub=document.getElementById("mySelect_sub");
	    if (val=="4"){
	    	mySelect_sub.style.display="inline-block";
	    }else{
	    	mySelect_sub.style.display="none";
	    }
     }
</script>
</body>
</html>