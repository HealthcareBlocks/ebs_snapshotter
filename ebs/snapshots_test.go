package ebs

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/stretchr/testify/assert"
)

// integration test that covers snapshot creation, tagging, and deletion
func TestSnapshotVolumes(t *testing.T) {
	mgr := NewSnapshotManager("us-west-1", awsServer.URL, true, 1, false)
	assert.EqualValues(t, mgr.SnapshotVolumes(), 1)
}

func TestSnapshotDestroyRemovesCorrectQuantity(t *testing.T) {
	volume := ec2.Volume{VolumeId: aws.String("vol-1a2b3c4d")}

	mgr := NewSnapshotManager("us-west-1", awsServer.URL, true, 1, false)
	assert.EqualValues(t, mgr.DestroySnapshots(&volume), 1)

	mgr = NewSnapshotManager("us-west-1", awsServer.URL, true, 2, false)
	assert.EqualValues(t, mgr.DestroySnapshots(&volume), 0)
}

// local test server
var awsServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	params := r.PostForm.Encode()

	if strings.Contains(params, "DescribeVolumes") {
		fmt.Fprintln(w, DescribeVolumesResponse)
	} else if strings.Contains(params, "CreateSnapshot") {
		fmt.Fprintln(w, CreateSnapshotResponse)
	} else if strings.Contains(params, "DescribeSnapshots") {
		fmt.Fprintln(w, DescribeSnapshotsResponse)
	} else if strings.Contains(params, "DeleteSnapshot") {
		fmt.Fprintln(w, DeleteSnapshotResponse)
	} else {
		fmt.Fprintln(w, "Response not implemented for this request type")
	}
}))

var DescribeVolumesResponse = `
<DescribeVolumesResponse xmlns="http://ec2.amazonaws.com/doc/2014-06-15/">
<requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
<volumeSet>
  <item>
	 <volumeId>vol-1a2b3c4d</volumeId>
	 <size>80</size>
	 <snapshotId/>
	 <availabilityZone>us-east-1a</availabilityZone>
	 <status>in-use</status>
	 <createTime>2013-12-18T22:35:00.000Z</createTime>
	 <attachmentSet>
		<item>
		   <volumeId>vol-1a2b3c4d</volumeId>
		   <instanceId>i-1a2b3c4d</instanceId>
		   <device>/dev/sdh</device>
		   <status>attached</status>
		   <attachTime>2013-12-18T22:35:00.000Z</attachTime>
		   <deleteOnTermination>false</deleteOnTermination>
		</item>
	 </attachmentSet>
	 <volumeType>standard</volumeType>
	 <encrypted>true</encrypted>
	 <tagSet>
	 	<item>
		   <key>aws:foo</key>
		   <value>bar</value>
		</item>
        <item>
           <key>Name</key>
           <value>Data Volume</value>
        </item>
		<item>
           <key>Mount Point</key>
           <value>/dev/xvdf</value>
        </item>
     </tagSet>
  </item>
</volumeSet>
</DescribeVolumesResponse>`

var CreateSnapshotResponse = `
<CreateSnapshotResponse xmlns="http://ec2.amazonaws.com/doc/2015-10-01/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <snapshotId>snap-2</snapshotId>
  <volumeId>vol-1a2b3c4d</volumeId>
  <status>pending</status>
  <startTime>2016-02-24T22:35:00.000Z</startTime>
  <progress>0%</progress>
  <ownerId>111122223333</ownerId>
  <volumeSize>15</volumeSize>
  <description>Daily Backup</description>
</CreateSnapshotResponse>
`

var DescribeSnapshotsResponse = `
<DescribeSnapshotsResponse xmlns="http://ec2.amazonaws.com/doc/2015-10-01/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <snapshotSet>
      <item>
         <snapshotId>snap-2</snapshotId>
         <volumeId>vol-1a2b3c4d</volumeId>
         <status>pending</status>
         <startTime>2016-02-24T22:35:00.000Z</startTime>
         <progress>30%</progress>
         <ownerId>111122223333</ownerId>
         <volumeSize>15</volumeSize>
         <description>Daily Backup</description>
         <encrypted>true</encrypted>
		 <tagSet/>
      </item>
	  <item>
         <snapshotId>snap-1</snapshotId>
         <volumeId>vol-1a2b3c4d</volumeId>
         <status>completed</status>
         <startTime>2016-02-23T22:35:00.000Z</startTime>
         <progress>100%</progress>
         <ownerId>111122223333</ownerId>
         <volumeSize>15</volumeSize>
         <description>Daily Backup</description>
         <encrypted>true</encrypted>
		 <tagSet/>
      </item>
   </snapshotSet>
</DescribeSnapshotsResponse>
`

var DeleteSnapshotResponse = `
<DeleteSnapshotResponse xmlns="http://ec2.amazonaws.com/doc/2015-10-01/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <return>true</return>
</DeleteSnapshotResponse>
`
