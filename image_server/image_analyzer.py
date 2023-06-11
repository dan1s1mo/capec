import cv2
import easyocr

def is_on_same_line(box1, box2, tolerance=10):
    y1 = (box1[0][1] + box1[2][1]) / 2
    y2 = (box2[0][1] + box2[2][1]) / 2

    return abs(y1 - y2) <= tolerance

def merge_boxes(box1, box2):
    x_coords = [point[0] for point in box1 + box2]
    y_coords = [point[1] for point in box1 + box2]
    return [(min(x_coords), min(y_coords)), (max(x_coords), max(y_coords))]

def concatenate_same_line_text(results):
    results.sort(key=lambda r: (r[0][0][1], r[0][0][0])) 
    lines = []
    line_text = []
    line_box = None

    for i, (box, text, _) in enumerate(results):
        line_text.append(text)

        if line_box is None:
            line_box = merge_boxes(box, box)
        else:
            line_box = merge_boxes(line_box, box)

        if i < len(results) - 1 and not is_on_same_line(box, results[i + 1][0]):
            lines.append((' '.join(line_text), line_box))
            line_text = []
            line_box = None

    if line_text:
        lines.append((' '.join(line_text), line_box))

    return lines

def transform_response(response): 
    text, box = response
    left, top = box[0]
    width = box[1][0] - left
    height = box[1][1] - top
    return {'text': text, 'left': int(left), 'top': int(top), 'width': int(width), 'height':int(height)}

def preprocess_image(image_path):
    image = cv2.imread(image_path)
    return image

def extract_text(image):
    reader = easyocr.Reader(['en'])
    result = reader.readtext(image, detail=1)
    lines = concatenate_same_line_text(result)
    return lines

def get_image_from_text(image_path):
    preprocessed_image = preprocess_image(image_path)
    text = extract_text(preprocessed_image)
    let = [transform_response(word) for word in text]
    print(let)
    return let 

