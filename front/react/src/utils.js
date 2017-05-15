import i18next from 'i18next';
import TimeAgo from 'timeago.js';


export const timeago = (d) => {
  var lang ;
  switch(i18next.language){
    case 'zh-Hans':
    case 'zh-CN':
      lang = 'zh_CN'
      break;
    case 'zh-Hant':
    case 'zh-TW':
      lang = 'zh_TW'
      break;
    default:
      lang = 'en_US'
  }

  return new TimeAgo(null, lang).format(d)
}
