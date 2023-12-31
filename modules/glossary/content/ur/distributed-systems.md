---
title: تقسیم شدہ نظام(Distributed System)
status: Completed
category: تصور
tags: ["فن تعمیر", "", ""]
---

## یہ کیا ہے

تقسیم شدہ نظام خود مختار کمپیوٹنگ عناصر کا ایک مجموعہ ہے جو ایک نیٹ ورک پر جڑا ہوتا ہے جو صارفین کو ایک اکٹھے نظام کے طور پر ظاہر ہوتا ہے۔ عام طور پر انہیں [نوڈز](/nodes/) کہا جاتا ہے، یہ اجزاء ہارڈویئر ڈیوائسز (مثلاً کمپیوٹر، موبائل فون) یا سافٹ ویئر ہو سکتے ہیں۔ نوڈز کو ایک مشترکہ مقصد حاصل کرنے کے لیے پروگرام کیا جاتا ہے ،اس کے علاوہ وہ تعاون کرنے کے لیے وہ نیٹ ورک پر پیغامات کا تبادلہ کرتے ہیں۔


## یہ کيا مسئلہ حل کرتا ہے

آج کل متعدد جدید ایپلی کیشنز اتنی بڑی ہیں کہ انہیں چلانے کے لیے سپر کمپیوٹرز کی ضرورت ہوگی۔ جی میل یا نیٹ فلکس پر غور کریں۔ کوئی ایک کمپیوٹر اتنا طاقتور نہیں ہے کہ پوری ایپلیکیشن کے یوزرز کی میزبانی کر سکے۔ متعدد کمپیوٹرز کو جوڑنے سے، کمپیوٹ کی طاقت تقریباً لامحدود ہو جاتی ہے۔ تقسیم شدہ کمپیوٹنگ کے بغیر، بہت سی ایپلیکیشنز جن پر ہم آج انحصار کرتے ہیں ممکن نہیں ہونگی۔

روایتی طور پر، سسٹمز عمودی 
[بڑھاؤ](/scalability/) کرتے تھے۔ یہ تب ہوتا ہے جب آپ انفرادی مشین میں مزید سی پی یو یا میموری شامل کرتے ہیں۔ عمودی اسکیلنگ وقت کی ضیاع ہے، اس کے لیے ڈاؤن ٹائم کی ضرورت ہوتی ہے، اور یہ تیزی سے اپنی حد تک پہنچ جاتی ہے۔

## یہ کس طرح مدد کرتا ہے

تقسیم شدہ نظام [افقی بڑھاؤ](/horizontal-scaling/) کی اجازت دیتے ہیں (مثال کے طور پر جب بھی ضرورت ہو سسٹم میں مزید نوڈز شامل کرنا)۔ یہ خود کار طریقے سے ہو سکتا ہے جس سے سسٹم کو کام کے بوجھ یا وسائل کی کھپت میں اچانک اضافے کو سنبھالا جا سکتا ہے۔
ایک غیر تقسیم شدہ نظام خود کو ناکامی کے خطرات سے دوچار کرتا ہے کیونکہ اگر ایک مشین ناکام ہوجاتی ہے تو پورا نظام ناکام ہوجاتا ہے۔ تقسیم شدہ نظام کو اس طرح سے ڈیزائن کیا جا سکتا ہے کہ اگر کچھ مشینیں  کام کرنا بند کر دیں تو بھی مجموعی نظام وہی نتائیج دینے کے لیے کام جاری رکھ سکے۔

