// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Battlefield) DeepCopyInto(out *Battlefield) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Battlefield.
func (in *Battlefield) DeepCopy() *Battlefield {
	if in == nil {
		return nil
	}
	out := new(Battlefield)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Battlefield) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BattlefieldList) DeepCopyInto(out *BattlefieldList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Battlefield, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BattlefieldList.
func (in *BattlefieldList) DeepCopy() *BattlefieldList {
	if in == nil {
		return nil
	}
	out := new(BattlefieldList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BattlefieldList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BattlefieldSpec) DeepCopyInto(out *BattlefieldSpec) {
	*out = *in
	if in.Players != nil {
		in, out := &in.Players, &out.Players
		*out = make([]Player, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BattlefieldSpec.
func (in *BattlefieldSpec) DeepCopy() *BattlefieldSpec {
	if in == nil {
		return nil
	}
	out := new(BattlefieldSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BattlefieldStatus) DeepCopyInto(out *BattlefieldStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.StopTime != nil {
		in, out := &in.StopTime, &out.StopTime
		*out = (*in).DeepCopy()
	}
	if in.Scores != nil {
		in, out := &in.Scores, &out.Scores
		*out = make([]PlayerStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BattlefieldStatus.
func (in *BattlefieldStatus) DeepCopy() *BattlefieldStatus {
	if in == nil {
		return nil
	}
	out := new(BattlefieldStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Player) DeepCopyInto(out *Player) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Player.
func (in *Player) DeepCopy() *Player {
	if in == nil {
		return nil
	}
	out := new(Player)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlayerStatus) DeepCopyInto(out *PlayerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlayerStatus.
func (in *PlayerStatus) DeepCopy() *PlayerStatus {
	if in == nil {
		return nil
	}
	out := new(PlayerStatus)
	in.DeepCopyInto(out)
	return out
}
