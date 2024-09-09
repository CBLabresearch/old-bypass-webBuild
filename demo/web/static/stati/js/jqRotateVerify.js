;(function(undefined) {
    "use strict"
    var _global;

    //判断是否为 数组
    function isArray(o){
        return Object.prototype.toString.call(o)=='[object Array]';
    }
    //随机数
    function getRandomNumber(a,b) {
        return Math.round(Math.random()*(b- a) + a)
    }
    //获取随机图片
    function getRandomImg(imgArr) {
        return imgArr[getRandomNumber(0,imgArr.length-1)]
    }
    //判断 是否处于动画状态
    function ifAnimate(ele) {
        if(ele.is(":animated")){
            return true
        }else{
            return false
        }
    }
    //获取元素的left值
    function getEleCssLeft($ele) {
        return parseInt($ele.css('left'));
    }
    
    var RotateVerify = function (ele,opt) {
        this.$ele = $(ele);
        //默认参数
        this.defaults = {
            initText:'滑动将图片转正',
            slideImage:'',
            slideAreaNum:10,
            getSucessState:function(){
				
            }
        }
        this.settings = $.extend({}, this.defaults, opt);
        this.init();
    }
    RotateVerify.prototype = {
        constructor: this,
        init:function () {
			this.verifyState = false;
			this.disLf = 0;
            this.initDom();
            this.initCanvasImg();
            this.initMouse();
            this._touchstart();
            this._touchend();
        },
        initDom:function () {
            var slideDom = $('<div class="rotateverify-contaniner">'+
					'<div class="rotate-can-wrap">'+
						'<canvas class="rotateCan rotate-can" width="200" height="200"></canvas>'+
						'<div class="statusBg status-bg"></div>'+
					'</div>'+
					'<div class="control-wrap slideDragWrap">'+
						'<div class="control-tips">'+
							'<p class="c-tips-txt cTipsTxt">'+ this.settings.initText +'</p>'+
						'</div>'+
						'<div class="control-bor-wrap controlBorWrap"></div>'+
						'<div class="control-btn slideDragBtn">'+
							'<i class="control-btn-ico"></i>'+
						'</div>'+
					'</div>'+
				'</div>');
            this.$ele.append(slideDom);
			this.statusBg = this.$ele.find('.statusBg');
            this.$slideDragWrap = this.$ele.find('.slideDragWrap');
            this.$slideDragBtn = this.$ele.find('.slideDragBtn');
			this.rotateCan = this.$ele.find('.rotateCan');
			this.cTipsTxt = this.$ele.find('.cTipsTxt');
			this.controlBorWrap = this.$ele.find('.controlBorWrap');
			this.xPos = this.rotateCan[0].width/2;
			this.yPos = this.rotateCan[0].height/2;
			this.aveRot = Math.round((360/(this.$slideDragWrap.width() - this.$slideDragBtn.outerWidth())) * 100 )/100;
			this.rotateImgCan = this.rotateCan[0].getContext('2d');
            this.slideImage = document.createElement('img');
			this.setAttrSrc();
        },
        initCanvasImg:function () {
			 this.randRot = getRandomNumber(30,270);
			 this.sucLenMin = (360 - this.settings.slideAreaNum - this.randRot) * (this.$slideDragWrap.width() - this.$slideDragBtn.width())/360;
			 this.sucLenMax = (360 + this.settings.slideAreaNum - this.randRot) * (this.$slideDragWrap.width() - this.$slideDragBtn.width())/360;
			 this.disLf = 0;
			 this.initImgSrc();
        },
		initImgSrc:function(){
			var _this = this;
			_this.slideImage.src = _this.slideImage.getAttribute('data-src');
			_this.setAttrSrc();
			_this.slideImage.onload = function(){
				_this.slideImage.style.width = '200px';
				_this.slideImage.style.height = '200px';
				_this.drawImgCan();
			}
		},
		drawImgCan:function(val){
			var _this = this;
			_this.rotateImgCan.beginPath();
			_this.rotateImgCan.arc( 100, 100, 100, 0, 360 * Math.PI / 180, false );
			_this.rotateImgCan.closePath();
			_this.rotateImgCan.clip();
			_this.rotateImgCan.save();
			_this.rotateImgCan.clearRect(0,0,200,200);
			_this.rotateImgCan.translate(_this.xPos, _this.yPos);
			_this.rotateImgCan.rotate(this.randRot * Math.PI / 180 + _this.disLf * _this.aveRot * Math.PI / 180);
			_this.rotateImgCan.translate(-_this.xPos, -_this.yPos);
			_this.rotateImgCan.drawImage( _this.slideImage, _this.xPos - 200 / 2, _this.yPos - 200 / 2,200,200);
			_this.rotateImgCan.restore();
		},
        initMouse:function () {
            var _this = this ;
            var ifThisMousedown = false;
            _this.$slideDragBtn.on('mousedown',function (e) {
                if(_this.verifyState){
                    return false;
                }
                if(_this.dragTimerState){
                    return false;
                }
                if(ifAnimate(_this.$slideDragBtn)){
                    return false;
                }
                ifThisMousedown = true;
                var positionDiv = $(this).offset();
                var distenceX = e.pageX - positionDiv.left;
                var disPageX = e.pageX;
                _this.$slideDragBtn.addClass('control-btn-active');
				_this.controlBorWrap.addClass('control-bor-active');
                $(document).mousemove(function (e) {
                    if(!ifThisMousedown){
                        return false;
                    }
                    var x = e.pageX - disPageX;
                    if(x<0){
                        x=0;
                    }else if(x >=(_this.$slideDragWrap.width()-_this.$slideDragBtn.outerWidth())){
						x = _this.$slideDragWrap.width()-_this.$slideDragBtn.outerWidth();
                    }
					_this.$slideDragBtn.css({
						'left':x + 'px'
					})
					_this.controlBorWrap.css({
						'width':x + _this.$slideDragBtn.width() + 'px'
					})
					_this.disLf = x;
					_this.drawImgCan();
                    e.preventDefault();
                })
            });
            $(document).on('mouseup',function(){
                if(!ifThisMousedown){
                    return false;
                }
                ifThisMousedown = false;
                if(_this.verifyState){
                    return false;
                }
                $(document).off('mousemove');
                _this.$slideDragBtn.removeClass('control-btn-active');
				_this.controlBorWrap.removeClass('control-bor-active');
                if(_this.sucLenMin <= _this.disLf && _this.disLf <= _this.sucLenMax){
                	_this.$slideDragBtn.addClass('control-btn-suc');
					_this.controlBorWrap.addClass('control-bor-suc');
					_this.statusBg.fadeIn();
					_this.statusBg.addClass('suc-bg');
					_this.verifyState = true;
					_this.cTipsTxt.text("");
					if(_this.settings.getSuccessState){
					    _this.settings.getSuccessState(_this.verifyState);
					}
                }else{
                	_this.$slideDragBtn.addClass('control-btn-err');
					_this.controlBorWrap.addClass('control-bor-err');
                	_this.$slideDragWrap.addClass('control-horizontal');
					_this.dragTimerState = true;
					_this.verifyState = false;;
					_this.statusBg.fadeIn();
					_this.statusBg.addClass('err-bg');
                	_this.$slideDragBtn.delay(700).animate({
                		left:'0px'
                	},function(){
						_this.dragTimerState = false;
                		_this.$slideDragWrap.removeClass('control-horizontal');
                		_this.$slideDragBtn.removeClass('control-btn-err');
						_this.statusBg.removeClass('err-bg');
						_this.statusBg.fadeOut();
						_this.refreshSlide();
                	})
					_this.controlBorWrap.delay(700).animate({
						width:_this.$slideDragBtn.width() + 'px'
					},function(){
						_this.controlBorWrap.removeClass('control-bor-err');
					});
                }
            });
        },
        _touchstart:function(){
            var _this = this;
			_this.$slideDragBtn.on('touchstart',function(e){
                _this.$slideDragBtn.css('pointer-events','none');
                setTimeout(function(){_this.$slideDragBtn.css('pointer-events','all')},400)
                if(_this.dragTimerState || ifAnimate(_this.$slideDragBtn) || _this.verifyState){
                    return false;
                }
                if(getEleCssLeft(_this.$slideDragBtn) == 0){
                    _this.touchX = e.originalEvent.targetTouches[0].pageX;
					_this.$slideDragBtn.addClass('control-btn-active');
					_this.controlBorWrap.addClass('control-bor-active');
                    _this._touchmove();
                }
            })
        },
        _touchmove:function(){
            var _this = this;
            _this.$slideDragBtn.on('touchmove',function(e){
                e.preventDefault();
                if(_this.dragTimerState || ifAnimate(_this.$slideDragBtn)){
                    return false;
                }else{
                    var x = e.originalEvent.targetTouches[0].pageX - _this.touchX;
					if(x<0){
					    x=0;
					}else if(x >=(_this.$slideDragWrap.width()-_this.$slideDragBtn.outerWidth())){
						x = _this.$slideDragWrap.width()-_this.$slideDragBtn.outerWidth();
					}
					_this.$slideDragBtn.css({
						'left':x + 'px'
					})
					_this.controlBorWrap.css({
						'width':x + _this.$slideDragBtn.width() + 'px'
					})
					_this.disLf = x;
					_this.drawImgCan();
                }
            })
        },
        _touchend:function(){
            var _this = this;
            _this.$slideDragBtn.on('touchend',function(){
                _this.$slideDragBtn.off('touchmove');
				_this.$slideDragBtn.removeClass('control-btn-active');
				_this.controlBorWrap.removeClass('control-bor-active');
				if((_this.sucLenMin) <= _this.disLf && _this.disLf <= (_this.sucLenMax)){
					_this.verifyState = true;
					_this.$slideDragBtn.addClass('control-btn-suc');
					_this.controlBorWrap.addClass('control-bor-suc');
					_this.statusBg.fadeIn();
					_this.statusBg.addClass('suc-bg');
					_this.cTipsTxt.text("");
					if(_this.settings.getSuccessState){
					    _this.settings.getSuccessState(_this.verifyState);
					}
				}else{
					if(!ifAnimate(_this.$slideDragBtn)){
					    _this.dragTimerState = true;
						_this.verifyState = false;
						_this.statusBg.fadeIn();
						_this.statusBg.addClass('err-bg');
					    _this.$slideDragBtn.addClass('control-btn-err');
						_this.controlBorWrap.addClass('control-bor-err');
					    _this.$slideDragWrap.addClass('control-horizontal');
					    _this.$slideDragBtn.delay(700).animate({
					    	left:'0px'
					    },function(){
					    	_this.$slideDragWrap.removeClass('control-horizontal');
					    	_this.$slideDragBtn.removeClass('control-btn-err');
							_this.statusBg.removeClass('err-bg');
							_this.statusBg.fadeOut();
					    	_this.dragTimerState = false;
					    	_this.refreshSlide();
					    })
						_this.controlBorWrap.delay(700).animate({
							width:_this.$slideDragBtn.width() + 'px'
						},function(){
							_this.controlBorWrap.removeClass('control-bor-err');
						});
					}else{
					    return false;
					}
				}
            })
        },
		setAttrSrc:function(){
			if(isArray(this.settings.slideImage)){
			    this.slideImageSrc = getRandomImg(this.settings.slideImage);
			}else{
			    this.slideImageSrc = this.settings.slideImage;
			}
			this.slideImage.setAttribute("data-src",this.slideImageSrc);
		},
        refreshSlide:function(){
            var _this = this;
			_this.initCanvasImg();
        },
		resetSlide:function(){
			var _this = this;
			_this.$slideDragBtn.css({
				left:'0px'
			})
			_this.controlBorWrap.css({
				width:_this.$slideDragBtn.width() + 'px'
			})
			_this.controlBorWrap.removeClass('control-bor-suc');
			_this.dragTimerState = false;
			_this.verifyState = false;;
			_this.$slideDragBtn.removeClass('control-btn-suc');
			_this.$slideDragWrap.removeClass('control-horizontal');
			_this.statusBg.fadeOut();
			_this.statusBg.removeClass('suc-bg');
			_this.cTipsTxt.text(_this.settings.initText);
			_this.refreshSlide();
		}
    }
	var inlineCss = '*{margin:0;padding:0;box-sizing:border-box;}.rotateverify-contaniner{width:200px;margin:0 auto;}@-webkit-keyframes rotateverifyHorizontal{0%{-webkit-transform:translate(0px,0);-ms-transform:translate(0px,0);transform:translate(0px,0)}10%,30%,50%,70%,90%{-webkit-transform:translate(-1px,0);transform:translate(-1px,0)}20%,40%,60%,80%{-webkit-transform:translate(1px,0);transform:translate(1px,0)}100%{-webkit-transform:translate(0px,0);transform:translate(0px,0)}}@-moz-keyframes rotateverifyHorizontal{0%{-webkit-transform:translate(0px,0);-moz-transform:translate(0px,0);transform:translate(0px,0)}10%,30%,50%,70%,90%{-webkit-transform:translate(-1px,0);-moz-transform:translate(-1px,0);transform:translate(-1px,0)}20%,40%,60%,80%{-webkit-transform:translate(1px,0);-moz-transform:translate(1px,0);transform:translate(1px,0)}100%{-webkit-transform:translate(0px,0);-moz-transform:translate(0px,0);transform:translate(0px,0)}}@keyframes rotateverifyHorizontal{0%{-webkit-transform:translate(0px,0);-moz-transform:translate(0px,0);transform:translate(0px,0)}10%,30%,50%,70%,90%{-webkit-transform:translate(-1px,0);-moz-transform:translate(-1px,0);transform:translate(-1px,0)}20%,40%,60%,80%{-webkit-transform:translate(1px,0);-moz-transform:translate(1px,0);transform:translate(1px,0)}100%{-webkit-transform:translate(0px,0);-moz-transform:translate(0px,0);transform:translate(0px,0)}}.rotateverify-contaniner .control-horizontal{-webkit-animation:rotateverifyHorizontal .6s .2s ease both;-moz-animation:rotateverifyHorizontal .6s .2s ease both;animation:rotateverifyHorizontal .6s .2s ease both}.rotateverify-contaniner .rotate-can-wrap{width:200px;height:200px;position:relative;}.rotateverify-contaniner .status-bg{width:100%;height:100%;position:absolute;top:0;left:0;background-color:rgba(0,0,0,.3);background-repeat:no-repeat;background-position:center center;border-radius:100%;display:none;}.rotateverify-contaniner .status-bg.suc-bg{background-image:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAH4AAABWCAYAAAAJ3CLTAAAAAXNSR0IArs4c6QAABFZJREFUeAHtncuO00AQRWdALFkiXsM0Az/FArGIhPg3EAvEazEL4AdY824eQsBXhNsTW1gmiduu6vct6cp20qmuuiftOMkoc3DAqM6B9Xq9gp5Dvzq5/VV1jbKhjQOAewSdQrvC3XeVflXkAIAeQ593ER/cflpR22230kH/MoA7tbs637Zl5XfvoKOL19DJjG4uEPwMt3IbuhC6a+PiudyaYT1+Dgign01A8H4+ZzVKCh3NvCX4rJBOF6MA3U3yeHomjsjGAQcdmnP1juH/Bd/OZUPUoxDg04DunjTXPabjkBwcUITu3voxSnCA0EugpFwjoSsbWkI6Qi+BknKNStDdFzZ8TVdmEyydIvQbwYpkYl0HCF3XzyKyEXoRmHSLJHRdP4vIRuhFYNItktB1/SwiG6EXgUm3SEA3kPRbNvc+nW/ZdNGEy9ZBt9hKgtDDIdLPDNIGspAkCF0fTbiMIG0gC0mC0MMh0s8M0gaykCQIXR9NuIwgbSALSYLQwyHSzwzSBrKQJAhdH024jCBtIAtJgtDDIdLPDNIGspAkCF0fTbiMIG0gC0mC0MMh0s8M0gaykCTiQ0e1d6CH0KdOj9xt+hbVlxE+GegrJAnne7yPYTHZJejJnordfZfrw6XTEbwxkIUkERe6ax3Vut9ImYp3GMCf0Rg9V+CJgSwkiSTQ786omPAH4OGbgSwkifjQu9X+dGbVhL85Sxr4Zmd6Nx6eBnoH/ue4Go/jpuHDHwNZSBIO+tHgBBJ3F5P/WVh9k/DhlYEsJIm00LsV/1LQQVPw4ZOBLCSJ9NA78A8kXeCx76Hqr/bRo4E03qenO70PX0zQzCH0BpJE1SsfxhjIQpLIY6WP4Bt09EPSFR5b5cpHXwaqZ6UPwXen/NtoUAq/qpUPP25WDb1/EqBJDfhVrPxmoBN+78DZR9htrPR/LW/2Wl75za10wm94pe+A/x2rQBJFvOajwTZP72Po/TEMcRd8VcMn9J72aFszfEIfwR4fwqBbkMbKvzbOneqY0D2drwk+oXtC74fVAJ/Qe5oztyXDV4L+EXny+JZtJjvx8BLhE7oY+yZBSfAJXQl6n6YE+ITe01Le5gyf0JVhj9PlCF8ROv+dxxj48Dgn+IQ+JBNhPwf4hB4B9LYpUsIn9G1EIt6mBP8D8nh/tk/oEQHvmyomfMx1Akn/GtZ9IscLuX1Qfe+LAb+D/g1bSRC6L1TfcaCh8ZXu1tM+cruVTui+MGKPCwGf0GNTXDifJnxCXwgh1cM6+NJTszvta+TghVzMJ4ISfKRZHLyQiwl8OBeQaVyULSFP6EMQKfYTwCf0FKC3zRkRvrsu4Gv6NgipbosAn9BTwZ2aNyB8Qp8yP/X9AeATemqovvMrwid0X9NzGacAn9BzgTm3DgF8Qp9rdm7jF8An9NwgLq1nBnxCX2pyro8D/GPoFbQr3M+yXsm1ftYlcABgD6H70AvoN+R+dfsZdE+QttmH/gVo1WKZD73PfwAAAABJRU5ErkJggg==);}.rotateverify-contaniner .status-bg.err-bg{background-image:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGoAAABqCAYAAABUIcSXAAAAAXNSR0IArs4c6QAABOhJREFUeAHt3F1O3DAUBeCZ7qJQUVXdQ6U+tUtBILq79okKUEGwAlaACO0DFV1C+0DPzXBnMpnYiR0nsZ1jyXUm/ok5H8wwBbJYVMrz8/MR6nfU3y9Vjo8qQ3gYOAHke4h6ivqIKrlL5seoy51L4eRr1B+opiJ9+zsTecI7AeS5h3phChznr2TM1gVwwoaka93j4GBrIh94JYAc36DeabCWVrBWX1k4kKe7roVYXjSbSQi6K5KanJSz8UieE10KsTa5Ox0hZFckcTlXKHkBcy3EciJaLBCwD5K4PMmlXjleT4e/w8ENFuFrliZiaQVJ8kJ9bxlm7RKoW+sIcyexzNmsewIglT4C9XW9qvsBsSyZBUCS1Tc+WLDLt+cYZix8zaqBISnf16RqyJtvz2V99OyjSth9CrFesBBiCKRC1qn5l1gH6CDWTjJuJwIivTVeGRchljGd9o5RkHQbxNIk3NpRkXRrxNIkurWTIOnWiKVJ2NtJkXRrxNIkmtsokHRrxNIkttuokHRrxNIkVm2USLpFYiWARKyEkOaOFfXTneLU27k9DSaJpGhzwUoaaS5YWSDljpUVUq5YWSLlhpU1Ui5Ys0BKHWtWSKlizRIpNaxZI6WCRSSVQoswovyFGSJVkPQwNiwiqUxDGwsWkRpw6qemxiJSXcTyeCosIllQTF1jYxHJJNHh/FhYROqA0TZkaCwitQk49A+FRSQHhK5DQ2MRqWvyHuMCYn3AWneofUqByea/T/L4+LKagnBC/HfT3z5CmFugEqntMwshhcDCMl6lwCwitSFpP8KaAotICuDSjoxFJBec+tiRsIhUD97n8cBYRPJBMc0ZCItIpsD7nA+MRaQ+GG1zgSVvZvu+T5L5H9uuxX7PBBBuiNvWYJmy3ONf3sLO08I4DaGGRFpRrW4nRCxj6o4dAyERy9HBOnxgJGJZ0+/YORISsTp6NA4bGYlYjQotJydCih7L9y7NLXH7dQsSZt6get/NGHP/ofqWaO+RGw1UIKQHCH1CLXylMC9arB4fU5ipgZ7uCqxT/tAPbYifZ/FNcZU3NJKuTSxNIkA7FJJujViaRI92aCTdGrE0CY92LCTdGrE0CYd2bCTdGrE0iQ7tVEi6NWJpEpZ2aiTdGrE0iYY2FiTdGrE0iUobG5JujViaBNpYkXSLxEoAiVgJIc0aK/anO8Wpt7N6GkwVSdFmgZU60iywckHKGis3pCyxckXKCit3pCyw5oKUNNbckJLEmitSUlhzR0oCi0jKtGqRR3y/N0ikbSR9FBUWkZSluY0Ci0jNOPWzk2IRqc5hfzwJFpHsKKbeUbGIZGLodn4ULCJ1w2gbNSgWFt9D5e092xQ69gfE2t+6JBa+QO1TCkzmnSMrqSKPEG+KL9dLYsFD1D6lwGQirRPdHCCXEFjH5YpY7BTVtxSYSKSNzc4R8umLdaZQj55KRNphaT7RE+tJVpW/il82L289+4Dez8vl8qd1FDvLBJDTL8kLtShPuP1T+gjUrdu8xQPGE8kxtB5YpY9AfXO45gPGEskhsOpQT6yVD54/l6hXqG2lwAB+41BN3vMYOXb9BuMaYzcvTXggb3htWNInt79hCZQA8txHvUQ1lWt0bL/hlWvjpHxlnaCeoz69VDmWcxvVQBvlMuvMj5Hv2Uvef9BK5l9QmXmKnyT/AS/pCbJI4duRAAAAAElFTkSuQmCC);}.rotateverify-contaniner .control-wrap{width:200px;position:relative;height:42px;border:1px solid #e0e0e0;clear:both;border-radius:42px;margin-top:5px;background-color:#f7f7f7;}.rotateverify-contaniner .control-wrap .control-tips{position:relative;width:100%;height:100%;}.rotateverify-contaniner .c-tips-txt{height:40px;line-height:40px;position:absolute;width:100%;text-align:center;top:0;left:0;overflow:hidden;white-space:nowrap;text-overflow:ellipsis;}.rotateverify-contaniner .control-btn{position:absolute;left:0;top:0;width:40px;height:40px;border-radius:40px;border:1px solid #e0e0e0;background-color:#fff;}.rotateverify-contaniner .control-bor-wrap{position:absolute;left:0;top:0;width:40;height:40px;border-radius:40px;border:1px solid transparent;}.rotateverify-contaniner .control-bor-active{border:1px solid #1a91ed;}.rotateverify-contaniner .control-bor-err{border:1px solid #e01116;}.rotateverify-contaniner .control-bor-suc{border:1px solid limegreen;}.rotateverify-contaniner .control-btn-active{background:#1a91ed;}.rotateverify-contaniner .control-btn-err{background:#e01116;}.rotateverify-contaniner .control-btn-suc{background-color:limegreen;}.rotateverify-contaniner .control-btn-ico{display:block;width:20px;height:20px;background:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADQAAAAyCAYAAAATIfj2AAAAAXNSR0IArs4c6QAABAdJREFUaAXtWUloU0EYbhaNTQWhItiIBz2kaYkHodWDBxEUa9OFIBHxoOnB3oSCgoqCKdijC3rpQTBCcYsQaUsE0eJBRIXiQZEuehJzUxCa1mIWvxFHpsPMvCUzh8oLlDfzzbxvvn+Zmb9JQ4P38TzgecDzgOeBVeQBn0prKpUKLC0tXa7VagNkns/nu93Y2Hgxl8tVVO9ZjZniJesGVYsvLi6OwJizdA7a54BtRv+PgRR3+jTFS3T4LcSk+XEYle7p6TnJ4w77pngtDRLqrFarN/r6+nYKB+sAdfAqI4Ro3JHoW1epVB5hL2yQjCthU7xkUaVBgUBgGAfBB5E6iNpeKpWyojErzBQvWTegWnxubu5XPB5/jlRIY15IMDcWjUZL8/PzrwRjUsgUL1lQeWxTRYlEIoWIPKR97lmGx/dNTEy85HDLrgleZYSoIkTgY2trazOM2k0x5knStguRHJuZmSkxuGXTBK9yD7GKIpHIGfRfsxhtw9DI8vLyvUwmY5uPvqub11aEyOLT09PV9vb2pxB/HN0wFcQ8txWLxSC8PsVglk3dvLb2EKuqu7u7C/0C/kTv1rCfEthPT9h37LR18dqOEBWFCHzCyUZSay/FmKcPETwUi8Xu4yT7weCWTV28jnOeKOvs7CT30zOJymYc8zlcumsl41JYB68rg7D5q+Fw+BiM+ipShyjtQgF6RTSmwnTwivaBas0VY6jn9pTL5RcAhVW73+8/Ojk5+WDFSzY69fA63kOsntnZ2S+4n0qIyEEWZ9pdbW1tOcz7zmCWzXp4XaUcqwgRuIrUy7MYbcPQ9ShiM7Tv5OmWt26DiEik1hCMqooEw6gdItwO5oZXi0E41a5DuJALhr63I140xw2vUISIXIahwDwNY5KicRizgIs2Ixqzwtzy/nennOsI4eLchA1PjmThkY3o3HRzZNfL68ogUlXj4ryLVNsiSh0Y8xYXL6nOHX108Lq6h5qamoZhzIBILYz5htNpfz6fd3T3EC4dvI73kK6qmHeGLl5HKdff378VERiDGKEjEJkRN/866OS1bdDg4OAa1G05pNpG3rt/+1MdHR2XJGNSWDev7T3U0tJyDcYcFilD1IqhUOjA6OjogmhchenmFaYOLwD5fQSYrGp2/a2PCV7LlEsmk1FE4BZvJO1j7Lybr7BM8Soj1NvbG0Y99QapFqcGcM/HhUJBWPZw81Z0TfGSRYS3PF0dlQDZ5DJjPuPeSNO5Tp6meIkGZcohnU5IhP4MBoMp/PDl6IsQymWK19IgKoB/ooI+NT4+/o7H6+3r4FVGCAKzvEh4N4tDQHpI8PMlfVO86j2EAvMCfmPFmfCvbssSTCLSNmyK17YAb6LnAc8DngdWrQd+A6CJYdIFu4L0AAAAAElFTkSuQmCC) no-repeat center center;background-size:100% 100%;margin:10px;}.rotateverify-contaniner .control-btn-active .control-btn-ico{background:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADQAAAAyCAYAAAATIfj2AAAAAXNSR0IArs4c6QAAAoRJREFUaAXtmb9OVFEQxjcQKGyspIPK2FvrA+gDAC8AL6C+gGIidvx5ABtqeABCDZSEFjtNLLTRwoLo+jvZJcBhvpnrLlPs5kzyhT3z5/vOzM1y757b6zVrE2gTaBNoE5iWCfT7/VnwAXwbonyeHbe/wgHunTfcF6KboLZPYWGQAGEKbyDb6yFcroxl62GxkwBhCq8jOQg5wr+JPQ0JREIWr5C7diP8ESj7TODhdXb3T9Sl8IY7QPgBOAfKDkISIwGyFF5D6q4L8Sfgp+oI/5u7VbEnizdWJgPxZaehS2LPOxFVSVm8lYy9RHzbaeorsQW70vdm8fqqRBGeA8dA2RGBmZCoSqAmhbeSsZeIL4LvQNl7u9L3QpbC66sOo4i/AH9FR8X/shNRlZTFW8nYS8TfioaK+wdYsit9L3UpvL4qUYRnwCFQdkpgPiSqEqhJ4a1k7CXij8AXoGzXrvS9kKXw+qrDKOLPQLkPKVvtRFQlQZbCW8nYS8RfqW7w/wKP7Urfm8Xrqw6jiO8DZXudSIwkCFN4DanbLoSXwB/R0dnt7O6rUXj/+84utrOFX3Gdi5ou7ixerc0UX4srU9zjfIdSeHUnRNhwyn+jLN6omeh+seMSiCDNpPAKuYEb0ZQ7ehav20wJIvwOKCtP46M+y6Xwug2x2el52qaZ6HfLhjsNEcziFXIDN6Lll+UJUHZEQN2LJDc1KbxS8CqA8PScKdDMClBWnrZHPfVJ4b26COZfNjs953I0k3LCmcVrXpGbToS9M+gL4hln2yPz3ty7+ZkNq9ceE/v2QTW0Zk6go9MZ1Fi8oTzCm6C2iX6DV78LLQ1mvGO9F97wCrWENoE2gTaBiZ/AP+8/LMb6T9MeAAAAAElFTkSuQmCC) no-repeat center center;background-size:100% 100%;}.rotateverify-contaniner .control-btn-err .control-btn-ico{background:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGoAAABqCAYAAABUIcSXAAAAAXNSR0IArs4c6QAABOhJREFUeAHt3F1O3DAUBeCZ7qJQUVXdQ6U+tUtBILq79okKUEGwAlaACO0DFV1C+0DPzXBnMpnYiR0nsZ1jyXUm/ok5H8wwBbJYVMrz8/MR6nfU3y9Vjo8qQ3gYOAHke4h6ivqIKrlL5seoy51L4eRr1B+opiJ9+zsTecI7AeS5h3phChznr2TM1gVwwoaka93j4GBrIh94JYAc36DeabCWVrBWX1k4kKe7roVYXjSbSQi6K5KanJSz8UieE10KsTa5Ox0hZFckcTlXKHkBcy3EciJaLBCwD5K4PMmlXjleT4e/w8ENFuFrliZiaQVJ8kJ9bxlm7RKoW+sIcyexzNmsewIglT4C9XW9qvsBsSyZBUCS1Tc+WLDLt+cYZix8zaqBISnf16RqyJtvz2V99OyjSth9CrFesBBiCKRC1qn5l1gH6CDWTjJuJwIivTVeGRchljGd9o5RkHQbxNIk3NpRkXRrxNIkurWTIOnWiKVJ2NtJkXRrxNIkmtsokHRrxNIkttuokHRrxNIkVm2USLpFYiWARKyEkOaOFfXTneLU27k9DSaJpGhzwUoaaS5YWSDljpUVUq5YWSLlhpU1Ui5Ys0BKHWtWSKlizRIpNaxZI6WCRSSVQoswovyFGSJVkPQwNiwiqUxDGwsWkRpw6qemxiJSXcTyeCosIllQTF1jYxHJJNHh/FhYROqA0TZkaCwitQk49A+FRSQHhK5DQ2MRqWvyHuMCYn3AWneofUqByea/T/L4+LKagnBC/HfT3z5CmFugEqntMwshhcDCMl6lwCwitSFpP8KaAotICuDSjoxFJBec+tiRsIhUD97n8cBYRPJBMc0ZCItIpsD7nA+MRaQ+GG1zgSVvZvu+T5L5H9uuxX7PBBBuiNvWYJmy3ONf3sLO08I4DaGGRFpRrW4nRCxj6o4dAyERy9HBOnxgJGJZ0+/YORISsTp6NA4bGYlYjQotJydCih7L9y7NLXH7dQsSZt6get/NGHP/ofqWaO+RGw1UIKQHCH1CLXylMC9arB4fU5ipgZ7uCqxT/tAPbYifZ/FNcZU3NJKuTSxNIkA7FJJujViaRI92aCTdGrE0CY92LCTdGrE0CYd2bCTdGrE0iQ7tVEi6NWJpEpZ2aiTdGrE0iYY2FiTdGrE0iUobG5JujViaBNpYkXSLxEoAiVgJIc0aK/anO8Wpt7N6GkwVSdFmgZU60iywckHKGis3pCyxckXKCit3pCyw5oKUNNbckJLEmitSUlhzR0oCi0jKtGqRR3y/N0ikbSR9FBUWkZSluY0Ci0jNOPWzk2IRqc5hfzwJFpHsKKbeUbGIZGLodn4ULCJ1w2gbNSgWFt9D5e092xQ69gfE2t+6JBa+QO1TCkzmnSMrqSKPEG+KL9dLYsFD1D6lwGQirRPdHCCXEFjH5YpY7BTVtxSYSKSNzc4R8umLdaZQj55KRNphaT7RE+tJVpW/il82L289+4Dez8vl8qd1FDvLBJDTL8kLtShPuP1T+gjUrdu8xQPGE8kxtB5YpY9AfXO45gPGEskhsOpQT6yVD54/l6hXqG2lwAB+41BN3vMYOXb9BuMaYzcvTXggb3htWNInt79hCZQA8txHvUQ1lWt0bL/hlWvjpHxlnaCeoz69VDmWcxvVQBvlMuvMj5Hv2Uvef9BK5l9QmXmKnyT/AS/pCbJI4duRAAAAAElFTkSuQmCC) no-repeat center center;background-size:100% 100%;}.rotateverify-contaniner .control-btn-suc .control-btn-ico{background:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAH4AAABWCAYAAAAJ3CLTAAAAAXNSR0IArs4c6QAABFZJREFUeAHtncuO00AQRWdALFkiXsM0Az/FArGIhPg3EAvEazEL4AdY824eQsBXhNsTW1gmiduu6vct6cp20qmuuiftOMkoc3DAqM6B9Xq9gp5Dvzq5/VV1jbKhjQOAewSdQrvC3XeVflXkAIAeQ593ER/cflpR22230kH/MoA7tbs637Zl5XfvoKOL19DJjG4uEPwMt3IbuhC6a+PiudyaYT1+Dgign01A8H4+ZzVKCh3NvCX4rJBOF6MA3U3yeHomjsjGAQcdmnP1juH/Bd/OZUPUoxDg04DunjTXPabjkBwcUITu3voxSnCA0EugpFwjoSsbWkI6Qi+BknKNStDdFzZ8TVdmEyydIvQbwYpkYl0HCF3XzyKyEXoRmHSLJHRdP4vIRuhFYNItktB1/SwiG6EXgUm3SEA3kPRbNvc+nW/ZdNGEy9ZBt9hKgtDDIdLPDNIGspAkCF0fTbiMIG0gC0mC0MMh0s8M0gaykCQIXR9NuIwgbSALSYLQwyHSzwzSBrKQJAhdH024jCBtIAtJgtDDIdLPDNIGspAkCF0fTbiMIG0gC0mC0MMh0s8M0gaykCTiQ0e1d6CH0KdOj9xt+hbVlxE+GegrJAnne7yPYTHZJejJnordfZfrw6XTEbwxkIUkERe6ax3Vut9ImYp3GMCf0Rg9V+CJgSwkiSTQ786omPAH4OGbgSwkifjQu9X+dGbVhL85Sxr4Zmd6Nx6eBnoH/ue4Go/jpuHDHwNZSBIO+tHgBBJ3F5P/WVh9k/DhlYEsJIm00LsV/1LQQVPw4ZOBLCSJ9NA78A8kXeCx76Hqr/bRo4E03qenO70PX0zQzCH0BpJE1SsfxhjIQpLIY6WP4Bt09EPSFR5b5cpHXwaqZ6UPwXen/NtoUAq/qpUPP25WDb1/EqBJDfhVrPxmoBN+78DZR9htrPR/LW/2Wl75za10wm94pe+A/x2rQBJFvOajwTZP72Po/TEMcRd8VcMn9J72aFszfEIfwR4fwqBbkMbKvzbOneqY0D2drwk+oXtC74fVAJ/Qe5oztyXDV4L+EXny+JZtJjvx8BLhE7oY+yZBSfAJXQl6n6YE+ITe01Le5gyf0JVhj9PlCF8ROv+dxxj48Dgn+IQ+JBNhPwf4hB4B9LYpUsIn9G1EIt6mBP8D8nh/tk/oEQHvmyomfMx1Akn/GtZ9IscLuX1Qfe+LAb+D/g1bSRC6L1TfcaCh8ZXu1tM+cruVTui+MGKPCwGf0GNTXDifJnxCXwgh1cM6+NJTszvta+TghVzMJ4ISfKRZHLyQiwl8OBeQaVyULSFP6EMQKfYTwCf0FKC3zRkRvrsu4Gv6NgipbosAn9BTwZ2aNyB8Qp8yP/X9AeATemqovvMrwid0X9NzGacAn9BzgTm3DgF8Qp9rdm7jF8An9NwgLq1nBnxCX2pyro8D/GPoFbQr3M+yXsm1ftYlcABgD6H70AvoN+R+dfsZdE+QttmH/gVo1WKZD73PfwAAAABJRU5ErkJggg==) no-repeat center center;background-size:100% 100%;}'
	var styleObj = $(
		'<style type="text/css">'+ inlineCss +'</style>'
	)
	$('head').prepend(styleObj);
    _global = (function(){ return this || (0, eval)('this'); }());
    if (typeof module !== "undefined" && module.exports) {
        module.exports = RotateVerify;
    } else if (typeof define === "function" && define.amd) {
        define(function(){return RotateVerify;});
    } else {
        !('RotateVerify' in _global) && (_global.RotateVerify = RotateVerify);
    }
}());