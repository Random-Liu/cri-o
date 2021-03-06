// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vision

import (
	"reflect"
	"testing"

	pb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

func TestImageToProtos(t *testing.T) {
	const url = "https://www.example.com/test.jpg"
	langHints := []string{"en", "fr"}
	img := NewImageFromURI("https://www.example.com/test.jpg")
	img.LanguageHints = langHints

	goti, gotc := img.toProtos()
	wanti := &pb.Image{Source: &pb.ImageSource{ImageUri: url}}
	if !reflect.DeepEqual(goti, wanti) {
		t.Errorf("got %+v, want %+v", goti, wanti)
	}
	wantc := &pb.ImageContext{
		LanguageHints: langHints,
	}
	if !reflect.DeepEqual(gotc, wantc) {
		t.Errorf("got %+v, want %+v", gotc, wantc)
	}
}
