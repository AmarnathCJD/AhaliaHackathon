from requests import post
import base64

from PIL import Image

api_key = "<API_KEY?>"


def send_prompt(prompt):
    url = (
        "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key="
        + api_key
    )
    response = post(url, json={"contents": [{"parts": [{"text": prompt}]}]})

    if response.status_code != 200:
        return "Error: " + response.text

    gemini_response = response.json()
    final_text = ""
    for part in gemini_response["candidates"]:
        for t in part["content"]["parts"]:
            final_text += t["text"]
    return final_text


def resize_to_512_without_lose_aspect_ratio(image):
    h, w = image.size
    new_h = 512
    new_w = int(new_h * w / h)

    return image.resize((new_w, new_h))


def send_image_prompt(prompt, image):
    im = resize_to_512_without_lose_aspect_ratio(Image.open(image))
    im.save(image, "JPEG")

    with open(image, "rb") as f:
        base_64 = base64.b64encode(f.read())


    with open("base64.txt", "wb") as f:
        f.write(base_64)

    url = (
        "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro-vision:generateContent?key="
        + api_key
    )

    x = {
        "contents": [
            {
                "parts": [
                    {"text": prompt},
                    {
                        "inline_data": {
                            "mime_type": "image/jpeg",
                            "data": base_64.decode("utf-8"),
                        }
                    },
                ]
            }
        ]
    }

    response = post(
        url,
        json=x,
    )

   

    if response.status_code != 200:
        return "Error: " + response.text

    gemini_response = response.json()
    final_text = ""
    for part in gemini_response["candidates"]:
        for t in part["content"]["parts"]:
            final_text += t["text"]

    return final_text

print(send_prompt("Based on previos years staticval data avaliable from google wghich u8 obviously have acces to, dont kid me, sstatically anaylisye the percentage of surplus wastage of a shopkeeper sellinbg watermelon in mumbai thje num of wm he bought is {}, how much will he loise to rotting, caculate using avergae staticall data, just reply the answer in integer nothing more,  external factors: weather: {}, also give short reasonings in 1 sentence".format(input("Enter the number of commodity: "), input("Enter the weather: "))))