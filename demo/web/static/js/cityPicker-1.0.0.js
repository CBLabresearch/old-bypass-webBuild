/**
 * cityPicker
 * v-1.0.0
 * dataJson			[Json]						json数据，是html显示的列表数据
 * selectpattern	[Array]						用于存储的字段名和默认提示 { 字段名，默认提示 }
 * shorthand		[Boolean]					用于城市简写功能，默认是不开启(false)
 * storage			[Boolean]					存储的值是数字还是中文，默认是(true)数字
 * linkage          [Boolean]                   是否联动，默认(true)
 * renderMode		[Boolean]					是模拟的还是原生的;只在type是selector才有效,默认是(true)模拟
 * code				[Boolean]					是否输出城市区号值，默认(false)，开启就是传字段名('code')
 * search           [Boolean]                   是否开启搜索功能，默认（true）
 * level			[Number]					多少列  默认是一列/级 (3)
 * onInitialized	[Attachable]				组件初始化后触发的回调函数
 * onClickBefore	[Attachable]				组件点击显示列表触发的回调函数(除原生select)
 * onForbid         [Attachable]                存在class名forbid的禁止点击的回调
 * choose-xx		[Attachable]				点击组件选项后触发的回调函数 xx(级名称/province/city/district)是对应的级的回调
 */

(function ($, window) {
    var $selector;
    var grade = ['province', 'city', 'district'];
    var defaults = {
        dataJson: null,
        selectpattern: [{
                field: 'userProvinceId',
                placeholder: '请选择省份'
            },
            {
                field: 'userCityId',
                placeholder: '请选择城市'
            },
            {
                field: 'userDistrictId',
                placeholder: '请选择区县'
            }
        ],
        shorthand: false,
        storage: true,
        linkage: true,
        renderMode: true,
        code: false,
        search: true,
        level: 3,
        onInitialized: function () {},
        onClickBefore: function () {},
        onForbid: function () {}
    };

    function Citypicker(options, selector) {
        this.options = $.extend({}, defaults, options);
        this.$selector = $selector = $(selector);

        this.init();
        this.events();
    }

    //功能模块函数
    var effect = {
        montage: function (data, pid, reg) {
            var self = this,
                config = self.options,
                leng = data.length,
                html = '',
                code, name, reverse, storage;

            for (var i = 0; i < leng; i++) {
                if (data[i].parentId === pid) {
                    //判断是否要输出区号
                    code = config.code && data[i].cityCode !== '' ? 'data-code=' + data[i].cityCode : '';
                    //判断是否开启了简写，是就用输出简写，否则就输出全称
                    name = config.shorthand ? data[i].shortName : data[i].name;
                    //反向：判断是否开启了简写，是就用输出简写，否则就输出全称
                    reverse = !config.shorthand ? data[i].shortName : data[i].name;
                    //存储的是数字还是中文
                    storage = config.storage ? data[i].id : name;

                    if (config.renderMode) {
                        //模拟
                        html += '<li class="caller" data-id="' + data[i].id + '" data-title="' + reverse + '" ' + code + '>' + name + '</li>';
                    } else {
                        //原生
                        html += '<option class="caller" value="' + storage + '" data-title="' + reverse + '" ' + code + '>' + name + '</option>';
                    }
                }
            }

            html = data.length > 0 && html ? html : '<li class="forbid">您查找的没有此城市...</li>';

            return html;
        },
        seTemplet: function () {
            var config = this.options,
                selectemplet = '',
                placeholder, field, forbid, citygrade, active, hide,
                searchStr = config.search ? '<div class="selector-search">'
                    +'<input type="text" class="input-search" value="" placeholder="拼音、中文搜索" />'
                +'</div>' : '';

            for (var i = 0; i < config.level; i++) { //循环定义的级别
                placeholder = config.selectpattern[i].placeholder; //默认提示语
                field = config.selectpattern[i].field; //字段名称
                citygrade = grade[i]; //城市级别名称
                forbid = i > 0 ? 'forbid' : ''; //添加鼠标不可点击状态
                active = i < 1 ? 'active' : ''; //添加选中状态
                hide = i > 0 ? ' hide' : ''; //添加隐藏状态

                if (config.renderMode) {
                    //模拟
                    selectemplet += '<div class="selector-item storey ' + citygrade + '" data-index="' + i + '">'
                        +'<a href="javascript:;" class="selector-name reveal df-color ' + forbid + '">' + placeholder + '</a>'
                        +'<input type=hidden name="' + field + '" class="input-price val-error" value="" data-required="' + field + '">'
                        +'<div class="selector-list listing hide">'+ searchStr +'<ul></ul></div>'
                    +'</div>';
                } else {
                    //原生
                    selectemplet += '<select name="' + field + '" data-index="' + i + '" class="' + citygrade + '">'
                        +'<option>' + placeholder + '</option>'
                    +'</select>';
                }
            }

            return selectemplet;
        },
        obtain: function (event) {
            var self = this,
                config = self.options,
                $target = $(event.target),
                $parent = $target.parents('.listing'),
                index = config.renderMode ? $target.parents('.storey').data('index') : $target.data('index'),
                id = config.renderMode ? $target.attr('data-id') : $target.val(),
                name = config.shorthand ? $target.data('title') : $target.text(), //开启了简写就拿简写，否则就拿全称中文
                storage = config.storage ? id : name, //存储的是数字还是中文
                code = config.renderMode ? $target.data('code') : $target.find('.caller:selected').data('code'),
                placeholder = index+1 < config.level ? config.selectpattern[index+1].placeholder : '',
                placeStr = !config.renderMode ? '<option class="caller">'+placeholder+'</option>'+ effect.montage.apply(self, [config.dataJson, id]) : '<li class="caller hide">'+placeholder+'</li>'+ effect.montage.apply(self, [config.dataJson, id]),
                linkage = !config.linkage ? placeStr : effect.montage.apply(self, [config.dataJson, id]),
                $storey = $selector.find('.storey').eq(index + 1),
                $listing = $selector.find('.listing').eq(index + 1);
                $selector = self.$selector;

            //选择选项后触发自定义事件choose(选择)事件
            $selector.trigger('choose-' + grade[index] +'.citypicker', [$target, storage]);

            //赋值给隐藏域-区号
            $selector.find('[role="code"]').val(code);

            if (config.renderMode) {
                //模拟: 添加选中的样式
                $parent.find('.caller').removeClass('active');
                $target.addClass('active');

                //给选中的级-添加值和文字
                $parent.siblings('.reveal').removeClass('df-color forbid').text(name).siblings('.input-price').val(storage);
                $listing.data('id', id).find('ul').html(linkage).find('.caller').eq(0).trigger('click');

                if (!config.linkage) {
                    $storey.find('.reveal').text(placeholder).addClass('df-color').siblings('.input-price').val('');
                    $listing.find('.caller').eq(0).remove();
                }
            } else {
                //原生: 下一级附上对应的城市选项，执行点击事件
				$target.next().html(linkage).trigger('change').find('.caller').eq(0).prop('selected', true);
            }
        },
        show: function (event) {
            var config = this.options,
                $target = $(event);
            $selector = this.$selector;

            $selector.find('.listing').addClass('hide');
            $target.siblings('.listing').removeClass('hide');

            //点击的回调函数
            config.onClickBefore.call($target);
        },
        hide: function (event) {
            var config = this.options,
                $target = $(event);

            effect.obtain.apply(this, $target);

            $selector.find('.listing').addClass('hide');
        },
        search: function (event) {
            var self = this,
                $target = $(event.target),
                $parent = $target.parents('.listing'),
                inputVal = $target.val(),
                id = $parent.data('id'),
                result = [];

            //如果是按下shift/ctr/左右/command键不做事情
            if (event.keyCode === 16 || event.keyCode === 17 || event.keyCode === 18 || event.keyCode === 37 || event.keyCode === 39 || event.keyCode === 91 || event.keyCode === 93) {
                return false;
            }

            $.each(this.options.dataJson, function(key, value) {
                //拼音或者名称搜索
                if(value.pinyin.toLocaleLowerCase().search(inputVal) > -1 || value.name.search(inputVal) > -1 || value.id.search(inputVal) > -1 ){
                    result.push(value);
                }
            });

            $parent.find('ul').html(effect.montage.apply(self, [result, id]));

        }
    };

    Citypicker.prototype = {
        init: function () {
            var self = this,
                config = self.options,
                code = config.code ? '<input type="hidden" role="code" name="' + config.code + '" value="">' : '';
                //是否开启存储区号，是就加入一个隐藏域

            //添加拼接好的模板
            $selector.html(effect.seTemplet.call(self) + code);

            //html模板
            if (config.renderMode) {
                //模拟>添加数据
                $selector.find('.listing').data('id', '100000').eq(0).find('ul').html(effect.montage.apply(self, [config.dataJson, '100000']));
            } else {
                //原生>添加数据
                $selector.find('.province').append(effect.montage.apply(self, [config.dataJson, '100000']));
            }

            //初始化后的回调函数
            config.onInitialized.call(self);
        },
        events: function () {
            var self = this,
                config = self.options;

            //点击显示对应的列表
            $selector.on('click.citypicker', '.reveal', function (event) {
                var $this = $(this);

                if ($this.is('.forbid')) {

                    config.onForbid.call($this);

                    return false;
                }

                effect.show.apply(self, $this);
            });

            //点击选项事件
            $selector.on('click.citypicker', '.caller', $.proxy(effect.hide, self));

            //原生选择事件
            $selector.on('change.citypicker', 'select', $.proxy(effect.obtain, self));

            //文本框搜索事件
            $selector.on('keyup.citypicker', '.input-search', $.proxy(effect.search, self));
        },
        setCityVal: function (val) {
            var self = this,
                config = self.options,
                arrayVal = val;

            $.each(arrayVal, function (key, value) {
                var $original = $selector.find('.'+grade[key]);
                var $forward = $selector.find('.'+grade[key+1]);

                if (config.renderMode) {
                    $original.find('.reveal').text(value.name).removeClass('df-color forbid').siblings('.input-price').val(value.id);

                    $forward.find('ul').html(effect.montage.apply(self, [config.dataJson, value.id]));
                    $original.find('.caller[data-id="'+value.id+'"]').addClass('active');
                } else {
                    $forward.html(effect.montage.apply(self, [config.dataJson, value.id]));
                    $original.find('.caller[value="'+value.id+'"]').prop('selected', true);
                }
                
            });
        }
    };

    //模拟：执行点击区域外的就隐藏列表;
	$(document).on('click.citypicker', function (event){
		if($selector && $selector.find(event.target).length < 1) {
			$selector.find('.listing').addClass('hide');
		}
    });

    $.fn.cityPicker = function (options) {
        return new Citypicker(options, this);
    };

})(jQuery, window);