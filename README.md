### filebeat-nsq

Collect log and then produce msg to nsq message queue.

### Usage

Modify **paths** and **output.nsq** columns in **examples/filebeat.yml** to your own special info.

Then exec  `make && ./build/filebeat --path.home examples`.
