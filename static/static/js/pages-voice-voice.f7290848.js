(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-voice-voice"],{"3b8d":function(t,e,r){"use strict";r.r(e),r.d(e,"default",function(){return a});var n=r("795b"),i=r.n(n);function o(t,e,r,n,o,a,c){try{var s=t[a](c),u=s.value}catch(l){return void r(l)}s.done?e(u):i.a.resolve(u).then(n,o)}function a(t){return function(){var e=this,r=arguments;return new i.a(function(n,i){var a=t.apply(e,r);function c(t){o(a,n,i,c,s,"next",t)}function s(t){o(a,n,i,c,s,"throw",t)}c(void 0)})}}},"3ff0":function(t,e,r){"use strict";r.r(e);var n=r("f24f"),i=r.n(n);for(var o in n)"default"!==o&&function(t){r.d(e,t,function(){return n[t]})}(o);e["default"]=i.a},"48e7":function(t,e,r){var n=r("be10");"string"===typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);var i=r("4f06").default;i("17f5e892",n,!0,{sourceMap:!1,shadowMode:!1})},"96cf":function(t,e){!function(e){"use strict";var r,n=Object.prototype,i=n.hasOwnProperty,o="function"===typeof Symbol?Symbol:{},a=o.iterator||"@@iterator",c=o.asyncIterator||"@@asyncIterator",s=o.toStringTag||"@@toStringTag",u="object"===typeof t,l=e.regeneratorRuntime;if(l)u&&(t.exports=l);else{l=e.regeneratorRuntime=u?t.exports:{},l.wrap=w;var f="suspendedStart",d="suspendedYield",p="executing",h="completed",v={},y={};y[a]=function(){return this};var g=Object.getPrototypeOf,m=g&&g(g(S([])));m&&m!==n&&i.call(m,a)&&(y=m);var b=k.prototype=T.prototype=Object.create(y);_.prototype=b.constructor=k,k.constructor=_,k[s]=_.displayName="GeneratorFunction",l.isGeneratorFunction=function(t){var e="function"===typeof t&&t.constructor;return!!e&&(e===_||"GeneratorFunction"===(e.displayName||e.name))},l.mark=function(t){return Object.setPrototypeOf?Object.setPrototypeOf(t,k):(t.__proto__=k,s in t||(t[s]="GeneratorFunction")),t.prototype=Object.create(b),t},l.awrap=function(t){return{__await:t}},E(L.prototype),L.prototype[c]=function(){return this},l.AsyncIterator=L,l.async=function(t,e,r,n){var i=new L(w(t,e,r,n));return l.isGeneratorFunction(e)?i:i.next().then(function(t){return t.done?t.value:i.next()})},E(b),b[s]="Generator",b[a]=function(){return this},b.toString=function(){return"[object Generator]"},l.keys=function(t){var e=[];for(var r in t)e.push(r);return e.reverse(),function r(){while(e.length){var n=e.pop();if(n in t)return r.value=n,r.done=!1,r}return r.done=!0,r}},l.values=S,P.prototype={constructor:P,reset:function(t){if(this.prev=0,this.next=0,this.sent=this._sent=r,this.done=!1,this.delegate=null,this.method="next",this.arg=r,this.tryEntries.forEach(I),!t)for(var e in this)"t"===e.charAt(0)&&i.call(this,e)&&!isNaN(+e.slice(1))&&(this[e]=r)},stop:function(){this.done=!0;var t=this.tryEntries[0],e=t.completion;if("throw"===e.type)throw e.arg;return this.rval},dispatchException:function(t){if(this.done)throw t;var e=this;function n(n,i){return c.type="throw",c.arg=t,e.next=n,i&&(e.method="next",e.arg=r),!!i}for(var o=this.tryEntries.length-1;o>=0;--o){var a=this.tryEntries[o],c=a.completion;if("root"===a.tryLoc)return n("end");if(a.tryLoc<=this.prev){var s=i.call(a,"catchLoc"),u=i.call(a,"finallyLoc");if(s&&u){if(this.prev<a.catchLoc)return n(a.catchLoc,!0);if(this.prev<a.finallyLoc)return n(a.finallyLoc)}else if(s){if(this.prev<a.catchLoc)return n(a.catchLoc,!0)}else{if(!u)throw new Error("try statement without catch or finally");if(this.prev<a.finallyLoc)return n(a.finallyLoc)}}}},abrupt:function(t,e){for(var r=this.tryEntries.length-1;r>=0;--r){var n=this.tryEntries[r];if(n.tryLoc<=this.prev&&i.call(n,"finallyLoc")&&this.prev<n.finallyLoc){var o=n;break}}o&&("break"===t||"continue"===t)&&o.tryLoc<=e&&e<=o.finallyLoc&&(o=null);var a=o?o.completion:{};return a.type=t,a.arg=e,o?(this.method="next",this.next=o.finallyLoc,v):this.complete(a)},complete:function(t,e){if("throw"===t.type)throw t.arg;return"break"===t.type||"continue"===t.type?this.next=t.arg:"return"===t.type?(this.rval=this.arg=t.arg,this.method="return",this.next="end"):"normal"===t.type&&e&&(this.next=e),v},finish:function(t){for(var e=this.tryEntries.length-1;e>=0;--e){var r=this.tryEntries[e];if(r.finallyLoc===t)return this.complete(r.completion,r.afterLoc),I(r),v}},catch:function(t){for(var e=this.tryEntries.length-1;e>=0;--e){var r=this.tryEntries[e];if(r.tryLoc===t){var n=r.completion;if("throw"===n.type){var i=n.arg;I(r)}return i}}throw new Error("illegal catch attempt")},delegateYield:function(t,e,n){return this.delegate={iterator:S(t),resultName:e,nextLoc:n},"next"===this.method&&(this.arg=r),v}}}function w(t,e,r,n){var i=e&&e.prototype instanceof T?e:T,o=Object.create(i.prototype),a=new P(n||[]);return o._invoke=C(t,r,a),o}function x(t,e,r){try{return{type:"normal",arg:t.call(e,r)}}catch(n){return{type:"throw",arg:n}}}function T(){}function _(){}function k(){}function E(t){["next","throw","return"].forEach(function(e){t[e]=function(t){return this._invoke(e,t)}})}function L(t){function e(r,n,o,a){var c=x(t[r],t,n);if("throw"!==c.type){var s=c.arg,u=s.value;return u&&"object"===typeof u&&i.call(u,"__await")?Promise.resolve(u.__await).then(function(t){e("next",t,o,a)},function(t){e("throw",t,o,a)}):Promise.resolve(u).then(function(t){s.value=t,o(s)},function(t){return e("throw",t,o,a)})}a(c.arg)}var r;function n(t,n){function i(){return new Promise(function(r,i){e(t,n,r,i)})}return r=r?r.then(i,i):i()}this._invoke=n}function C(t,e,r){var n=f;return function(i,o){if(n===p)throw new Error("Generator is already running");if(n===h){if("throw"===i)throw o;return N()}r.method=i,r.arg=o;while(1){var a=r.delegate;if(a){var c=R(a,r);if(c){if(c===v)continue;return c}}if("next"===r.method)r.sent=r._sent=r.arg;else if("throw"===r.method){if(n===f)throw n=h,r.arg;r.dispatchException(r.arg)}else"return"===r.method&&r.abrupt("return",r.arg);n=p;var s=x(t,e,r);if("normal"===s.type){if(n=r.done?h:d,s.arg===v)continue;return{value:s.arg,done:r.done}}"throw"===s.type&&(n=h,r.method="throw",r.arg=s.arg)}}}function R(t,e){var n=t.iterator[e.method];if(n===r){if(e.delegate=null,"throw"===e.method){if(t.iterator.return&&(e.method="return",e.arg=r,R(t,e),"throw"===e.method))return v;e.method="throw",e.arg=new TypeError("The iterator does not provide a 'throw' method")}return v}var i=x(n,t.iterator,e.arg);if("throw"===i.type)return e.method="throw",e.arg=i.arg,e.delegate=null,v;var o=i.arg;return o?o.done?(e[t.resultName]=o.value,e.next=t.nextLoc,"return"!==e.method&&(e.method="next",e.arg=r),e.delegate=null,v):o:(e.method="throw",e.arg=new TypeError("iterator result is not an object"),e.delegate=null,v)}function j(t){var e={tryLoc:t[0]};1 in t&&(e.catchLoc=t[1]),2 in t&&(e.finallyLoc=t[2],e.afterLoc=t[3]),this.tryEntries.push(e)}function I(t){var e=t.completion||{};e.type="normal",delete e.arg,t.completion=e}function P(t){this.tryEntries=[{tryLoc:"root"}],t.forEach(j,this),this.reset(!0)}function S(t){if(t){var e=t[a];if(e)return e.call(t);if("function"===typeof t.next)return t;if(!isNaN(t.length)){var n=-1,o=function e(){while(++n<t.length)if(i.call(t,n))return e.value=t[n],e.done=!1,e;return e.value=r,e.done=!0,e};return o.next=o}}return{next:N}}function N(){return{value:r,done:!0}}}(function(){return this||"object"===typeof self&&self}()||Function("return this")())},"9cda":function(t,e,r){"use strict";r.r(e);var n=r("c0ce"),i=r("3ff0");for(var o in i)"default"!==o&&function(t){r.d(e,t,function(){return i[t]})}(o);r("c154");var a=r("2877"),c=Object(a["a"])(i["default"],n["a"],n["b"],!1,null,"008757f3",null);e["default"]=c.exports},be10:function(t,e,r){e=t.exports=r("2350")(!1),e.push([t.i,"uni-image[data-v-008757f3]{width:%?150?%;height:%?150?%}.page-body-wrapper[data-v-008757f3]{-webkit-box-pack:justify;-webkit-justify-content:space-between;justify-content:space-between;-webkit-box-flex:1;-webkit-flex-grow:1;flex-grow:1;margin-bottom:%?300?%}.page-body-time[data-v-008757f3]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-orient:vertical;-webkit-box-direction:normal;-webkit-flex-direction:column;flex-direction:column;-webkit-box-align:center;-webkit-align-items:center;align-items:center}.time-big[data-v-008757f3]{font-size:%?60?%;margin:%?20?%}.time-small[data-v-008757f3]{font-size:%?30?%}.page-body-buttons[data-v-008757f3]{margin-top:%?60?%;display:-webkit-box;display:-webkit-flex;display:flex;-webkit-justify-content:space-around;justify-content:space-around}.page-body-button[data-v-008757f3]{width:%?250?%;text-align:center}.button-stop-record[data-v-008757f3]{width:%?110?%;height:%?110?%;border:%?20?% solid #fff;background-color:#f55c23;border-radius:%?130?%;margin:0 auto}",""])},c0ce:function(t,e,r){"use strict";var n=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("v-uni-view",[r("v-uni-view",{staticClass:"uni-padding-wrap"},[t.recording||t.playing||t.hasRecord?t._e():[r("v-uni-view",{staticClass:"page-body-time"},[r("v-uni-text",{staticClass:"time-big"},[t._v(t._s(t.formatedRecordTime))])],1),r("v-uni-view",{staticClass:"page-body-buttons"},[r("v-uni-view",{staticClass:"page-body-button"}),r("v-uni-view",{staticClass:"page-body-button",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.startRecord.apply(void 0,arguments)}}},[r("v-uni-image",{attrs:{src:"../../static/voice/record.png"}})],1),r("v-uni-view",{staticClass:"page-body-button"})],1)],!0===t.recording?[r("v-uni-view",{staticClass:"page-body-time"},[r("v-uni-text",{staticClass:"time-big"},[t._v(t._s(t.formatedRecordTime))])],1),r("v-uni-view",{staticClass:"page-body-buttons"},[r("v-uni-view",{staticClass:"page-body-button"}),r("v-uni-view",{staticClass:"page-body-button",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.stopRecord.apply(void 0,arguments)}}},[r("v-uni-view",{staticClass:"button-stop-record"})],1),r("v-uni-view",{staticClass:"page-body-button"})],1)]:t._e(),!0===t.hasRecord&&!1===t.playing?[r("v-uni-view",{staticClass:"page-body-time"},[r("v-uni-text",{staticClass:"time-big"},[t._v(t._s(t.formatedPlayTime))]),r("v-uni-text",{staticClass:"time-small"},[t._v(t._s(t.formatedRecordTime))])],1),r("v-uni-view",{staticClass:"page-body-buttons"},[r("v-uni-view",{staticClass:"page-body-button",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.playVoice.apply(void 0,arguments)}}},[r("v-uni-image",{attrs:{src:"../../static/voice/play.png"}})],1),r("v-uni-view",{staticClass:"page-body-button",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.clear.apply(void 0,arguments)}}},[r("v-uni-image",{attrs:{src:"../../static/voice/record.png"}})],1)],1)]:t._e(),!0===t.hasRecord&&!0===t.playing?[r("v-uni-view",{staticClass:"page-body-time"},[r("v-uni-text",{staticClass:"time-big"},[t._v(t._s(t.formatedPlayTime))]),r("v-uni-text",{staticClass:"time-small"},[t._v(t._s(t.formatedRecordTime))])],1),r("v-uni-view",{staticClass:"page-body-buttons"},[r("v-uni-view",{staticClass:"page-body-button",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.stopVoice.apply(void 0,arguments)}}},[r("v-uni-image",{attrs:{src:"../../static/voice/stop.png"}})],1),r("v-uni-view",{staticClass:"page-body-button",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.clear.apply(void 0,arguments)}}},[r("v-uni-image",{attrs:{src:"../../static/voice/trash.png"}})],1)],1)]:t._e()],2)],1)},i=[];r.d(e,"a",function(){return n}),r.d(e,"b",function(){return i})},c154:function(t,e,r){"use strict";var n=r("48e7"),i=r.n(n);i.a},c9cf:function(t,e,r){"use strict";var n=r("288e"),i=n(r("0a0d"));r("28a5");var o=n(r("59ad"));r("6b54");var a=n(r("e814"));function c(t){if("number"!==typeof t||t<0)return t;var e=(0,a.default)(t/3600);t%=3600;var r=(0,a.default)(t/60);t%=60;var n=t;return[e,r,n].map(function(t){return t=t.toString(),t[1]?t:"0"+t}).join(":")}function s(t,e){return"string"===typeof t&&"string"===typeof e&&(t=(0,o.default)(t),e=(0,o.default)(e)),t=t.toFixed(2),e=e.toFixed(2),{longitude:t.toString().split("."),latitude:e.toString().split(".")}}var u={UNITS:{"年":315576e5,"月":26298e5,"天":864e5,"小时":36e5,"分钟":6e4,"秒":1e3},humanize:function(t){var e="";for(var r in this.UNITS)if(t>=this.UNITS[r]){e=Math.floor(t/this.UNITS[r])+r+"前";break}return e||"刚刚"},format:function(t){var e=this.parse(t),r=(0,i.default)()-e.getTime();if(r<this.UNITS["天"])return this.humanize(r);var n=function(t){return t<10?"0"+t:t};return e.getFullYear()+"/"+n(e.getMonth()+1)+"/"+n(e.getDate())+"-"+n(e.getHours())+":"+n(e.getMinutes())},parse:function(t){var e=t.split(/[^0-9]/);return new Date(e[0],e[1]-1,e[2],e[3],e[4],e[5])}};t.exports={formatTime:c,formatLocation:s,dateUtils:u}},f24f:function(t,e,r){"use strict";var n=r("288e");Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0,r("96cf");var i=n(r("3b8d")),o=r("c9cf"),a=null,c=null,s=null,u=null,l={data:function(){return{string:"helloworld",title:"start/stopRecord、play/stopVoice",recording:!1,playing:!1,hasRecord:!1,tempFilePath:"",recordTime:0,playTime:0,formatedRecordTime:"00:00:00",formatedPlayTime:"00:00:00"}},onUnload:function(){this.end()},onLoad:function(){var t=this;u=uni.createInnerAudioContext(),u.onEnded(function(){clearInterval(a);var e=0;console.log("play voice finished"),t.playing=!1,t.formatedPlayTime=o.formatTime(e),t.playTime=e}),s=uni.getRecorderManager(),s.onStart(function(){console.log("recorder start"),t.recording=!0,c=setInterval(function(){t.recordTime+=1,t.formatedRecordTime=o.formatTime(t.recordTime)},1e3)}),s.onStop(function(e){console.log("on stop"),u.src=e.tempFilePath,t.hasRecord=!0,t.recording=!1}),s.onError(function(){console.log("recorder onError")})},methods:{startRecord:function(){var t=(0,i.default)(regeneratorRuntime.mark(function t(){return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:s.start({format:"mp3"});case 1:case"end":return t.stop()}},t,this)}));function e(){return t.apply(this,arguments)}return e}(),stopRecord:function(){s.stop(),clearInterval(c)},playVoice:function(){var t=this;console.log("play voice"),this.playing=!0,a=setInterval(function(){t.playTime+=1,t.formatedPlayTime=o.formatTime(t.playTime)},1e3),u.play()},stopVoice:function(){clearInterval(a),this.playing=!1,this.formatedPlayTime=o.formatTime(0),this.playTime=0,u.stop()},end:function(){u.stop(),s.stop(),clearInterval(c),clearInterval(a),this.recording=!1,this.playing=!1,this.hasRecord=!1,this.playTime=0,this.recordTime=0,this.formatedRecordTime="00:00:00",this.formatedRecordTime="00:00:00"},clear:function(){this.end()}}};e.default=l}}]);