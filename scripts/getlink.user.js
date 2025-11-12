// ==UserScript==
// @name         Gosuslugi Download Interceptor
// @namespace    http://tampermonkey.net/
// @version      0.2
// @description  Intercepts certificate download clicks on gosuslugi.ru/crt
// @author       zarazaex
// @match        https://www.gosuslugi.ru/crt*
// @grant        none
// @run-at       document-start
// ==/UserScript==

(function() {
    'use strict';

    window.addEventListener('click', (event) => {
        const button = event.target.closest('button');

        if (button && (button.textContent.includes('Скачать корневые сертификаты') || button.textContent.includes('Скачать выпускающие сертификаты'))) {
            console.warn('Download click intercepted. Default action prevented. Inspect the element or network tab.', button);
            event.preventDefault();
            event.stopPropagation();
        }
    }, true);

})();c
