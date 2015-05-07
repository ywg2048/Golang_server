<!DOCTYPE html>
<html lang="zh-cn">
<head>
<meta charset="UTF-8" />
<title>奇妙的朋友手游官网</title>
<meta name="viewport" content="width=device-width,initial-scale=1, minimum-scale=1.0, maximum-scale=1, user-scalable=no"/>
<meta name="apple-mobile-web-app-status-bar-style" content="black" />
<meta name="apple-mobile-web-app-capable" content="yes" />
<meta http-equiv="keywords" content="奇妙的朋友,手游,官网" />
<meta http-equiv="description" content="奇妙的朋友,手游,官网" />
<link rel="stylesheet" href="/static/css/style.css"/>
<script type="text/javascript" src="/static/js/jquery-1.9.1.js"></script>
<script type="text/javascript" src="/static/js/TouchSwipe-Jquery-Plugin.js"></script>
<script type="text/javascript" src="/static/js/alert.js"></script>
<script>
	var current_page = 1;
	var target = 'http://mm.10086.cn/download/android/300008817370?from=www';
	//var target = 'http://m.anzhi.com/info_2083716.html';
	var ua = navigator.userAgent;
    var wx = ua.match(/MicroMessenger\/(.+)/i);
	$(function(){
		$(document.body).swipe( {
			//Generic swipe handler for all directions
			swipe:function(event, direction, distance, duration, fingerCount, fingerData) {
				if(direction == 'up'){
					nextPage(4);
				}
				
				if(direction == 'down'){
					prevPage(4);
				}
			},
			//Default is 75px, set to 0 for demo so any distance triggers swipe
		   threshold:50,
		   fingers:'all'
		});
		
		for(var i = 1; i <=3;i++){
			$('#page_' + i).find('img').each(function(){
				$(this).width($(document).width() + "px");
				$(this).height($(document).height() + "px");
			});
		}
	});
	
	function nextPage(totalPage){
		if(current_page != totalPage){
			var next_page = current_page + 1 ;
			for(var i = 1;i<= totalPage; i ++){
				$('#page_' + i).slideUp(800);
			}
			$('#page_' + next_page).slideDown(800);
			current_page = next_page;
		}
	}
	
	function prevPage(totalPage){
		if(current_page != 1){
			var prev_page = current_page - 1 ;
			for(var i = 1;i<= totalPage; i ++){
				$('#page_' + i).slideUp(1000);
			}
			$('#page_' + prev_page).slideDown(1000);
			current_page = prev_page;
		}
	}
	function down(){
	    var wx = ua.match(/MicroMessenger\/(.+)/i);
        var qq5 = ua.match(/QQ\/5/i);
	
        //微信中打开
      
		  var osVersion = getOS();
            var ios = osVersion.indexOf("iOS") != -1;
            if (ios) {
               alert('即将上线，敬请期待');
                return;
            }
          if (wx) $('#share').show();
          window.location.href = target;

	}
	 function getOS() {
            //ios or android or unknow
            var userOS;
            // use Number(userOSver) to convert
            var userOSver;
            //get os version

            var uaindex;

            // determine OS
            if (ua.match(/iPad/i) || ua.match(/iPhone/i)) {
                userOS = 'iOS';
                uaindex = ua.indexOf('OS ');
            }
            else if (ua.match(/Android/i)) {
                userOS = 'Android';
                uaindex = ua.indexOf('Android ');
            }
            else {
                userOS = 'unknown';
            }

            // determine version
            if (userOS === 'iOS' && uaindex > -1) {
                userOSver = ua.substr(uaindex + 3, 3).replace('_', '.');
                return "iOS" + userOSver;
            }
            else if (userOS === 'Android' && uaindex > -1) {
                userOSver = ua.substr(uaindex + 8, 3);
                return "Android" + userOSver;
            }
            else {
                return "unKnow";
            }
        }
        
</script>
</head>
<body class="app-body" >

	<div class="page-wrap" id="page_1" style="background-color:#88BB0F;max-width:none;min-width:none;" >
		<div class="download">
			<img src="/static/img/bg.jpg" width="100%" height="100%" />
			<div class="cont">
				<!-- <a href="#">下载</a>-->
				<span class="bullet moveIconDown"><i></i></span>
			</div>
		</div>
	</div>
	<div class="page-wrap" id="page_2" style="background-color:#88BB0F;display:none;max-width:none;min-width:none;">
		<div class="download">
			<img src="/static/img/2.jpg" />
			<div class="cont">
				<span class="bullet moveIconDown"><i></i></span>
			</div>
		</div>
	</div>
	<div class="page-wrap" id="page_3" style="background-color:#88BB0F;display:none;max-width:none;min-width:none;">
		<div class="download">
			<img src="/static/img/3.jpg" />
			<div class="cont">
				<span class="bullet moveIconDown"><i></i></span>
			</div>
		</div>
	</div>
	<div class="page-wrap" id="page_4" style="display:none;">
		<div class="qrcode">
			<div class="pic"><img src="/static/img/qrcode.png" /></div>
			<p>『扫描二维码下载』</p>
		</div>
		<div class="app">
			<h1 class="app-title">奇妙的朋友</h1>
			<div class="app-info">
			 	<p>当前版本：1.1.0(Build 1.1.0)</p>
			 	<p>应用类型：游戏</p>
			 	<p>文件大小：28.35MB</p>
			 	<p><i class="icon-ios"></i><span class="lab lab-1" onclick="down();">安卓下载</span><span class="lab lab-2" onclick="down();">iPhone下载</span></p>
			</div>		
			<div class="app-intro">
				<h2 class="title">&gt;介绍</h2>			
				<div class="text">湖南卫视官方唯一授权正版《奇妙的朋友》手游。
国内首款人与动物互动的三消类手机游戏，春春、妮妮、轩轩、杏儿、涛涛、皓皓六位明星饲养员和COCO,六毛等宠物明星陪你共享消除类游戏最完美的体验。</div>
			</div>
		</div>
	</div>
	
		<div id="share" style="display: none">
			<img width="100%"  height="100%" src="/static/img/share.png" style="position: fixed; z-index: 9999; top: 0; left: 0; display: " ontouchstart="document.getElementById(&#39;share&#39;).style.display=&#39;none&#39;;" />
		</div>
<script>
(function(){
	window.addEventListener("resize",init);
	function init(){
		var w = window.innerWidth;
		w = (w >= 640) ? 640 : w;
		document.documentElement.style.fontSize = 16/320*w+"px";				
	}
	setTimeout(init,100);
})();	
</script>
</body>
</html>