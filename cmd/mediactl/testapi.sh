cat > /tmp/foo <<EOF
<?xml version="1.0" encoding="utf-8"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
<s:Body>
<u:GetVolume xmlns:u="urn:schemas-upnp-org:service:RenderingControl:1">
<InstanceID>0</InstanceID>
<Channel>Master</Channel>
</u:GetVolume>
</s:Body>
</s:Envelope>
EOF

curl -v -X POST --data-binary @/tmp/foo -H "SoapAction: \"urn:schemas-upnp-org:service:RenderingControl:1#GetVolume\"" 192.168.0.199:9197/upnp/control/RenderingControl1
