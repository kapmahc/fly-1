import Vue from 'vue'

import I18n from 'vue-i18n'
Vue.use(I18n)

import {get} from './ajax'

export const LOCALE = 'locale'

export const i18n = new I18n({})

export function detect () {
  // TODO detect from request
  return 'zh-Hans'
}

export function load (locale) {
  get(`/locales/${locale}`).then((message) => {
    i18n.setLocaleMessage(locale, message)
    i18n.locale = locale
    document.title = i18n.t('site.title')
  }).catch((err) => {
    console.error(err)
  })
}
