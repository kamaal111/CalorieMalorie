import time
import json
import xml.etree.cElementTree as ET


def main():
    start = time.time()
    xmlFile = ET.parse('DataSet/myfoodapediadata/Food_Display_Table.xml')
    root = xmlFile.getroot()

    food_display_table = []

    for food_display_row in root:
        food_item_dict = {}
        for food_item in food_display_row:
            value = None
            if food_item.tag == 'Food_Code':
                value = int(food_item.text)
            elif food_item.tag in ['Display_Name', 'Portion_Display_Name']:
                value = food_item.text
            else:
                value = float(food_item.text)
            
            key = None
            if food_item.tag == 'Drkgreen_Vegetables':
                key = 'darkgreen_vegetables'
            else:
                key = food_item.tag.lower()
            food_item_dict[key] = value
        food_display_table.append(food_item_dict)

    with open("somepy.json", 'w') as json_file:
        json.dump(food_display_table, json_file, indent=2, ensure_ascii=False)

    end = time.time()
    print(f"converted in {end - start}s")


if __name__ == "__main__":
    main()