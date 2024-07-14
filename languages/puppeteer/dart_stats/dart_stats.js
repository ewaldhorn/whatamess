const puppeteer = require('puppeteer');
const alphabet = 'poiuytrewqlkjhgfdsamnbvcxz'.split('');
(async () => {
  let results = {};
  const browser = await puppeteer.launch({args: ['--no-sandbox', '--disable-setuid-sandbox']});
  const page = await browser.newPage();
  await page.setViewport({ width: 1024, height: 768 });
  
  for (let char of alphabet) {
    await page.goto(`https://pub.dev/packages?q=package%3A${char}`);
    const el = await page.$('.package-count span');
    const count = await (await el.getProperty('textContent')).jsonValue();
    results[char] = parseInt(count);
    await page.screenshot({path: `pub/${char}.png`});
  }
  console.log(JSON.stringify(results, null, 2));
  await browser.close();
})();
