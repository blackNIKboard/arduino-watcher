#include <Wire.h>
#include <LiquidCrystal_I2C.h>

LiquidCrystal_I2C lcd(0x3F, 16, 2);

String line = "";
int rd = 0;

void setup()
{
	lcd.begin();
	lcd.backlight();

	Serial.begin(9600);
}

void loop() {
  while (Serial.available()) {
    char c = (char) Serial.read();

  if (c != '\n' && c != '`'){
    line += c;
    lcd.setCursor(0,1);
    lcd.write(c);
    rd++;
    continue;
  }
      String lines[2];

      lcd.clear();
      for (int j = 0; j < 2; j++){
        for (int i = 0; i < 16; i++){
          int strPos = i+(j*16);
          char cur = line.charAt(strPos);
          if (strPos >= rd) {
            cur = ' ';
          }
          lines[j]+=cur;
        }
        lcd.setCursor(0,j);
        lcd.print(lines[j]);
      }

      rd = 0;
      line = "";
  }
}