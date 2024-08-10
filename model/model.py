# model/model.py
import sys
import tensorflow as tf
from tensorflow.keras.applications.mobilenet_v2 import MobileNetV2, preprocess_input, decode_predictions
from tensorflow.keras.preprocessing import image
import numpy as np

def recognize_image(image_path):
    model = MobileNetV2(weights='imagenet')

    img = image.load_img(image_path, target_size=(224, 224))
    img_array = image.img_to_array(img)
    img_array = np.expand_dims(img_array, axis=0)
    img_array = preprocess_input(img_array)

    predictions = model.predict(img_array)
    decoded_predictions = decode_predictions(predictions, top=3)[0]
    
    return decoded_predictions

if __name__ == "__main__":
    image_path = sys.argv[1]
    results = recognize_image(image_path)
    for result in results:
        print(f"{result[1]}: {result[2]*100:.2f}%")
