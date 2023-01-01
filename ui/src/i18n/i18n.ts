import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';

import LanguageDetector from 'i18next-browser-languagedetector';

import en from './locales/en/translations.json';
import de from './locales/de/translations.json';

const fallbackLng = 'en';

i18n
  .use(LanguageDetector)
  .use(initReactI18next)
  .init({
    fallbackLng,
    debug: !import.meta.env.PROD,
    supportedLngs: ['de', 'en'],
    resources: {
      de: {
        translation: de,
      },
      en: {
        translation: en,
      },
    },

    interpolation: {
      escapeValue: false,
    },
    react: {
      useSuspense: false,
    },
  });

export default i18n;
