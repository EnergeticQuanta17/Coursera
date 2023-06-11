#! python3
import pyautogui
import time

time.sleep(2)

pyautogui.click(x=40, y=30)
for i in range(12):
	time.sleep(0.5)
	pyautogui.click(x=121, y=25)
