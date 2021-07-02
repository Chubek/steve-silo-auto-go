from selenium import webdriver
from dotenv import dotenv_values
from lxml.html.soupparser import fromstring

temp = dotenv_values(".env")

options = webdriver.ChromeOptions()
options.add_argument("--headless")
options.add_argument("--disable-extensions")
options.add_argument("--disable-gpu")
options.add_argument("--disable-dev-shm-usage")
options.add_argument("--no-sandbox")
prefs = {"profile.managed_default_content_settings.images": 2, "network.cookie.cookieBehavior": 2}
options.add_experimental_option("prefs", prefs)
options.add_experimental_option('excludeSwitches', ['enable-logging'])
driver = webdriver.Chrome(executable_path=temp["DRIVER_PATH"], chrome_options=options)

driver.get("https://us2.m2web.talk2m.com/valleycarriers/Gorman%20Bros/usr/viewon/Overview.shtm")

driver.find_element_by_xpath("//*[@id='username']").send_keys("operator")
driver.find_element_by_xpath("//*[@id='password']").send_keys("operator123")
driver.save_screenshot("s1.png")
driver.find_element_by_xpath("//*[@id='connect']").click()


driver.save_screenshot("ss.png")