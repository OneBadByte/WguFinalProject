from selenium import webdriver
from selenium.webdriver.common.keys import Keys


driver = webdriver.Firefox()
driver.get("abyss:8080/ping")
sleep(10)
driver.close()
