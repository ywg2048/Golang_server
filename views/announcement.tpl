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
	<a href="/stage">返回管理首页</a>
<table class="bordered">
	<tr><th>序号</th><th>标题</th><th>内容</th><th>标题2</th><th>内容2</th><th>是否激活</th><th>操作</th></tr>
	{{range $k, $v := .message}}
		<form action='/message' method='get'>
		<tr><td>{{$v.Id}}</td><td>{{$v.Title}}</td><td>{{$v.Content}}</td><td>{{$v.Title2}}</td><td>{{$v.Content2}}</td>
			<td>
				{{if eq $v.IsActive 0}}
				<font color="red">未激活</font>
				{{else}}
				<font color="green">已激活</font>
				{{end}}
			</td>
			<td>
				<a href="?id={{$v.Id}}&opration=active">激活</a><a href="?id={{$v.Id}}&opration=close">关闭</a>
			</td>
		</tr>
		</form>
	{{end}}

</table>

</center>

</body>
</html>