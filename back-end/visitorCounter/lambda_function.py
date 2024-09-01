import json
import boto3

dynamodb = boto3.resource('dynamodb')
table = dynamodb.Table('VisitorCount')

def lambda_handler(event, context):
    # Increment the counter
    response = table.update_item(
        Key={'id': 'counter'},
        UpdateExpression="ADD #count :incr",
        ExpressionAttributeNames={'#count': 'count'},
        ExpressionAttributeValues={':incr': 1},
        ReturnValues="UPDATED_NEW"
    )

    # Fetch the updated count
    new_count = response['Attributes']['count']

    return {
        'statusCode': 200,
        'body': json.dumps({'count': new_count})
    }
