# Get open/collection buckets (CAR files)
In this guide, we will get the open buckets or collections from the edge node.

## Pre-requisites
- make sure you have a edge node running either locally or remote. Use this guide [running a node](running_node.md) to run a node.
- identify the edge node host.
- get a API key using this guide [getting an API key](getting-api-key.md)
- collection name.


## Get open buckets
```
curl --location 'http://localhost:1313/buckets/get/open' \
--header 'Authorization: Bearer [ANY VALID DELTA API KEY]'
[
    {
        "bucket_uuid": "7f02a270-0f8f-11ee-b4e2-9e0bf0c70138",
        "piece_cid": "baga6ea4seaqgrg7bn7nycekq35wzt265pmdviyju7ae3psomwe6mlenykgu2wfi",
        "piece_size": 4194304,
        "download_url": "/gw/bafybeicqff43zsaz4njvvr3m2pu5pe7yyj5bfogud5ijmsmynstcts3y34",
        "status": "ready",
        "size": 2500485,
        "created_at": "2023-06-20T13:25:48.068312-04:00",
        "updated_at": "2023-06-20T13:29:49.23759-04:00"
    },
    {
        "bucket_uuid": "0ed5ec36-0f90-11ee-8024-9e0bf0c70138",
        "piece_cid": "baga6ea4seaqdmvydl5fwqy67zjeytacabatupzeu65e4qhsz2xnpveglvsikwky",
        "piece_size": 4194304,
        "download_url": "/gw/bafybeigbevcu26ec33qdgcgfquzbyeeebsx6ifjqjgnsiweu3yxhznmgei",
        "status": "ready",
        "collection_name": "mytag1",
        "size": 2500485,
        "created_at": "2023-06-20T13:29:49.36727-04:00",
        "updated_at": "2023-06-20T13:29:50.324344-04:00"
    },
    {
        "bucket_uuid": "0f7c1368-0f90-11ee-8024-9e0bf0c70138",
        "piece_cid": "baga6ea4seaqfhto4gcpkxy2tch2tnoljvzs4pgteuzztz467xzd4aqomoudfihq",
        "piece_size": 4194304,
        "download_url": "/gw/bafybeicdomb2q3sz4jvjeub6wuxfaxwmrehiv56zeh43trb3udg2mhdk54",
        "status": "ready",
        "collection_name": "mytag2",
        "size": 2500485,
        "created_at": "2023-06-20T13:29:50.456101-04:00",
        "updated_at": "2023-06-20T13:29:51.372687-04:00"
    },
    {
        "bucket_uuid": "101d1c0e-0f90-11ee-8024-9e0bf0c70138",
        "piece_cid": "baga6ea4seaqo24mj4c23wtvn4beyaedn6o7mkbv2fpflo3hq5tgtjn455w7m2my",
        "piece_size": 4194304,
        "download_url": "/gw/bafybeig2p6dlwzenrn2a4yo5izpqitmhxi2s2gwx75iitc4omcz44liej4",
        "status": "ready",
        "collection_name": "mytag3",
        "size": 2500485,
        "created_at": "2023-06-20T13:29:51.511452-04:00",
        "updated_at": "2023-06-20T13:35:54.900899-04:00"
    },
    {
        "bucket_uuid": "e8caad78-0f90-11ee-9f1c-9e0bf0c70138",
        "piece_cid": "baga6ea4seaqp2tctcdwnp3ed4prtvok65ievb3xa7ch4ib5v2uv2p6yzazrqedq",
        "piece_size": 4194304,
        "download_url": "/gw/bafybeib442ldouoxtd6abjk7m32e4pyvs4bdaat23jz6bnh2cia325ew5a",
        "status": "ready",
        "collection_name": "mytag1",
        "size": 2500485,
        "created_at": "2023-06-20T13:35:55.036883-04:00",
        "updated_at": "2023-06-20T13:35:55.708092-04:00"
    },
    {
        "bucket_uuid": "e9456054-0f90-11ee-9f1c-9e0bf0c70138",
        "piece_cid": "baga6ea4seaqcb3pxczxuvjj62j4lhsittepkbsnjvdzaujonsobaswg4dekracq",
        "piece_size": 4194304,
        "download_url": "/gw/bafybeieqnenofuqcz2van4gccihs723v6pi5p5viwx43tnetti2rdx5pre",
        "status": "ready",
        "collection_name": "mytag1",
        "size": 2500485,
        "created_at": "2023-06-20T13:35:55.84093-04:00",
        "updated_at": "2023-06-20T13:35:56.677806-04:00"
    },
    {
        "bucket_uuid": "e9dc0d6a-0f90-11ee-9f1c-9e0bf0c70138",
        "piece_cid": "baga6ea4seaqpv5jywdub7y4zsmevhwbs6mufx3nvgfjaic3ge7wdd3eto5c6iiq",
        "piece_size": 4194304,
        "download_url": "/gw/bafybeicumrvtfe52nteln7gifoblfp6bqauwu75ycuur4gvvs3doztzw3m",
        "status": "ready",
        "collection_name": "mytag1",
        "size": 2500485,
        "created_at": "2023-06-20T13:35:56.828402-04:00",
        "updated_at": "2023-06-20T14:03:53.399393-04:00"
    },
    {
        "bucket_uuid": "d166d31a-0f94-11ee-b379-9e0bf0c70138",
        "piece_cid": "baga6ea4seaqewl5llxjucsmqlev6qvljmingogd55n7dqvmmczzkszdr3xz6woi",
        "piece_size": 4194304,
        "download_url": "/gw/bafybeihgq7rccpg5alb5tavypeykugtbjpu6qdtsxuhwzf2x3v7uuep73q",
        "status": "ready",
        "collection_name": "mytag4",
        "size": 2500485,
        "created_at": "2023-06-20T14:03:53.781874-04:00",
        "updated_at": "2023-06-20T14:03:55.055813-04:00"
    }
]
```

## Get open buckets based on collection name
```
curl --location 'http://localhost:1313/collections/get?name=mytag1' \
--header 'Authorization: Bearer [ANY VALID DELTA API KEY]'
[
    {
        "bucket_uuid": "7f02a270-0f8f-11ee-b4e2-9e0bf0c70138",
        "piece_cid": "baga6ea4seaqgrg7bn7nycekq35wzt265pmdviyju7ae3psomwe6mlenykgu2wfi",
        "piece_size": 4194304,
        "download_url": "/gw/bafybeicqff43zsaz4njvvr3m2pu5pe7yyj5bfogud5ijmsmynstcts3y34",
        "status": "ready",
        "collection_name": "mytag1",
        "size": 2500485,
        "created_at": "2023-06-20T13:25:48.068312-04:00",
        "updated_at": "2023-06-20T13:29:49.23759-04:00"
    },
    {
        "bucket_uuid": "0ed5ec36-0f90-11ee-8024-9e0bf0c70138",
        "piece_cid": "baga6ea4seaqdmvydl5fwqy67zjeytacabatupzeu65e4qhsz2xnpveglvsikwky",
        "piece_size": 4194304,
        "download_url": "/gw/bafybeigbevcu26ec33qdgcgfquzbyeeebsx6ifjqjgnsiweu3yxhznmgei",
        "status": "ready",
        "collection_name": "mytag1",
        "size": 2500485,
        "created_at": "2023-06-20T13:29:49.36727-04:00",
        "updated_at": "2023-06-20T13:29:50.324344-04:00"
    },
    {
        "bucket_uuid": "0f7c1368-0f90-11ee-8024-9e0bf0c70138",
        "piece_cid": "baga6ea4seaqfhto4gcpkxy2tch2tnoljvzs4pgteuzztz467xzd4aqomoudfihq",
        "piece_size": 4194304,
        "download_url": "/gw/bafybeicdomb2q3sz4jvjeub6wuxfaxwmrehiv56zeh43trb3udg2mhdk54",
        "status": "ready",
        "collection_name": "mytag1",
        "size": 2500485,
        "created_at": "2023-06-20T13:29:50.456101-04:00",
        "updated_at": "2023-06-20T13:29:51.372687-04:00"
    },
    {
        "bucket_uuid": "101d1c0e-0f90-11ee-8024-9e0bf0c70138",
        "piece_cid": "baga6ea4seaqo24mj4c23wtvn4beyaedn6o7mkbv2fpflo3hq5tgtjn455w7m2my",
        "piece_size": 4194304,
        "download_url": "/gw/bafybeig2p6dlwzenrn2a4yo5izpqitmhxi2s2gwx75iitc4omcz44liej4",
        "status": "ready",
        "collection_name": "mytag1",
        "size": 2500485,
        "created_at": "2023-06-20T13:29:51.511452-04:00",
        "updated_at": "2023-06-20T13:35:54.900899-04:00"
    },
    {
        "bucket_uuid": "e8caad78-0f90-11ee-9f1c-9e0bf0c70138",
        "piece_cid": "baga6ea4seaqp2tctcdwnp3ed4prtvok65ievb3xa7ch4ib5v2uv2p6yzazrqedq",
        "piece_size": 4194304,
        "download_url": "/gw/bafybeib442ldouoxtd6abjk7m32e4pyvs4bdaat23jz6bnh2cia325ew5a",
        "status": "ready",
        "collection_name": "mytag1",
        "size": 2500485,
        "created_at": "2023-06-20T13:35:55.036883-04:00",
        "updated_at": "2023-06-20T13:35:55.708092-04:00"
    },
    {
        "bucket_uuid": "e9456054-0f90-11ee-9f1c-9e0bf0c70138",
        "piece_cid": "baga6ea4seaqcb3pxczxuvjj62j4lhsittepkbsnjvdzaujonsobaswg4dekracq",
        "piece_size": 4194304,
        "download_url": "/gw/bafybeieqnenofuqcz2van4gccihs723v6pi5p5viwx43tnetti2rdx5pre",
        "status": "ready",
        "size": 2500485,
        "created_at": "2023-06-20T13:35:55.84093-04:00",
        "updated_at": "2023-06-20T13:35:56.677806-04:00"
    },
    {
        "bucket_uuid": "e9dc0d6a-0f90-11ee-9f1c-9e0bf0c70138",
        "piece_cid": "baga6ea4seaqpv5jywdub7y4zsmevhwbs6mufx3nvgfjaic3ge7wdd3eto5c6iiq",
        "piece_size": 4194304,
        "download_url": "/gw/bafybeicumrvtfe52nteln7gifoblfp6bqauwu75ycuur4gvvs3doztzw3m",
        "status": "ready",
        "collection_name": "mytag1",
        "size": 2500485,
        "created_at": "2023-06-20T13:35:56.828402-04:00",
        "updated_at": "2023-06-20T14:03:53.399393-04:00"
    },
    {
        "bucket_uuid": "d166d31a-0f94-11ee-b379-9e0bf0c70138",
        "piece_cid": "baga6ea4seaqewl5llxjucsmqlev6qvljmingogd55n7dqvmmczzkszdr3xz6woi",
        "piece_size": 4194304,
        "download_url": "/gw/bafybeihgq7rccpg5alb5tavypeykugtbjpu6qdtsxuhwzf2x3v7uuep73q",
        "status": "ready",
        "collection_name": "mytag1",
        "size": 2500485,
        "created_at": "2023-06-20T14:03:53.781874-04:00",
        "updated_at": "2023-06-20T14:03:55.055813-04:00"
    }
]
```