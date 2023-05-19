from selenium import webdriver
import time

# replace with your own details
contact_name = "John Doe"
message = "Hello, this is a test message!"

# open WhatsApp Web in Chrome browser
driver = webdriver.Chrome()
driver.get("https://web.whatsapp.com/")

# wait for user to scan QR code and load the chat page
input("Please scan the QR code and press Enter to continue...")

# find the search box and type the contact name
search_box = driver.find_element_by_xpath("//input[@title='Search or start new chat']")
search_box.send_keys(contact_name)
search_box.submit()

# wait for chat to load and send the message
time.sleep(2)
message_box = driver.find_element_by_xpath("//div[@contenteditable='true'][@data-tab='1']")
message_box.send_keys(message)
send_button = driver.find_element_by_xpath("//span[@data-testid='send']")
send_button.click()

# close the browser
driver.quit()
