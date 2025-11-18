// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cjavdev/spotted-go/internal/apijson"
	"github.com/cjavdev/spotted-go/internal/requestconfig"
	"github.com/cjavdev/spotted-go/option"
	"github.com/cjavdev/spotted-go/packages/respjson"
)

// AudioAnalysisService contains methods and other services that help with
// interacting with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioAnalysisService] method instead.
type AudioAnalysisService struct {
	Options []option.RequestOption
}

// NewAudioAnalysisService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAudioAnalysisService(opts ...option.RequestOption) (r AudioAnalysisService) {
	r = AudioAnalysisService{}
	r.Options = opts
	return
}

// Get a low-level audio analysis for a track in the Spotify catalog. The audio
// analysis describes the track’s structure and musical content, including rhythm,
// pitch, and timbre.
//
// Deprecated: deprecated
func (r *AudioAnalysisService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AudioAnalysisGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("audio-analysis/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TimeIntervalObject struct {
	// The confidence, from 0.0 to 1.0, of the reliability of the interval.
	Confidence float64 `json:"confidence"`
	// The duration (in seconds) of the time interval.
	Duration float64 `json:"duration"`
	// The starting point (in seconds) of the time interval.
	Start float64 `json:"start"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		Duration    respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimeIntervalObject) RawJSON() string { return r.JSON.raw }
func (r *TimeIntervalObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioAnalysisGetResponse struct {
	// The time intervals of the bars throughout the track. A bar (or measure) is a
	// segment of time defined as a given number of beats.
	Bars []TimeIntervalObject `json:"bars"`
	// The time intervals of beats throughout the track. A beat is the basic time unit
	// of a piece of music; for example, each tick of a metronome. Beats are typically
	// multiples of tatums.
	Beats []TimeIntervalObject         `json:"beats"`
	Meta  AudioAnalysisGetResponseMeta `json:"meta"`
	// Sections are defined by large variations in rhythm or timbre, e.g. chorus,
	// verse, bridge, guitar solo, etc. Each section contains its own descriptions of
	// tempo, key, mode, time_signature, and loudness.
	Sections []AudioAnalysisGetResponseSection `json:"sections"`
	// Each segment contains a roughly conisistent sound throughout its duration.
	Segments []AudioAnalysisGetResponseSegment `json:"segments"`
	// A tatum represents the lowest regular pulse train that a listener intuitively
	// infers from the timing of perceived musical events (segments).
	Tatums []TimeIntervalObject          `json:"tatums"`
	Track  AudioAnalysisGetResponseTrack `json:"track"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bars        respjson.Field
		Beats       respjson.Field
		Meta        respjson.Field
		Sections    respjson.Field
		Segments    respjson.Field
		Tatums      respjson.Field
		Track       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioAnalysisGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AudioAnalysisGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioAnalysisGetResponseMeta struct {
	// The amount of time taken to analyze this track.
	AnalysisTime float64 `json:"analysis_time"`
	// The version of the Analyzer used to analyze this track.
	AnalyzerVersion string `json:"analyzer_version"`
	// A detailed status code for this track. If analysis data is missing, this code
	// may explain why.
	DetailedStatus string `json:"detailed_status"`
	// The method used to read the track's audio data.
	InputProcess string `json:"input_process"`
	// The platform used to read the track's audio data.
	Platform string `json:"platform"`
	// The return code of the analyzer process. 0 if successful, 1 if any errors
	// occurred.
	StatusCode int64 `json:"status_code"`
	// The Unix timestamp (in seconds) at which this track was analyzed.
	Timestamp int64 `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnalysisTime    respjson.Field
		AnalyzerVersion respjson.Field
		DetailedStatus  respjson.Field
		InputProcess    respjson.Field
		Platform        respjson.Field
		StatusCode      respjson.Field
		Timestamp       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioAnalysisGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *AudioAnalysisGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioAnalysisGetResponseSection struct {
	// The confidence, from 0.0 to 1.0, of the reliability of the section's
	// "designation".
	Confidence float64 `json:"confidence"`
	// The duration (in seconds) of the section.
	Duration float64 `json:"duration"`
	// The estimated overall key of the section. The values in this field ranging from
	// 0 to 11 mapping to pitches using standard Pitch Class notation (E.g. 0 = C, 1 =
	// C♯/D♭, 2 = D, and so on). If no key was detected, the value is -1.
	Key int64 `json:"key"`
	// The confidence, from 0.0 to 1.0, of the reliability of the key. Songs with many
	// key changes may correspond to low values in this field.
	KeyConfidence float64 `json:"key_confidence"`
	// The overall loudness of the section in decibels (dB). Loudness values are useful
	// for comparing relative loudness of sections within tracks.
	Loudness float64 `json:"loudness"`
	// Indicates the modality (major or minor) of a section, the type of scale from
	// which its melodic content is derived. This field will contain a 0 for "minor", a
	// 1 for "major", or a -1 for no result. Note that the major key (e.g. C major)
	// could more likely be confused with the minor key at 3 semitones lower (e.g. A
	// minor) as both keys carry the same pitches.
	//
	// Any of -1, 0, 1.
	Mode float64 `json:"mode"`
	// The confidence, from 0.0 to 1.0, of the reliability of the `mode`.
	ModeConfidence float64 `json:"mode_confidence"`
	// The starting point (in seconds) of the section.
	Start float64 `json:"start"`
	// The overall estimated tempo of the section in beats per minute (BPM). In musical
	// terminology, tempo is the speed or pace of a given piece and derives directly
	// from the average beat duration.
	Tempo float64 `json:"tempo"`
	// The confidence, from 0.0 to 1.0, of the reliability of the tempo. Some tracks
	// contain tempo changes or sounds which don't contain tempo (like pure speech)
	// which would correspond to a low value in this field.
	TempoConfidence float64 `json:"tempo_confidence"`
	// An estimated time signature. The time signature (meter) is a notational
	// convention to specify how many beats are in each bar (or measure). The time
	// signature ranges from 3 to 7 indicating time signatures of "3/4", to "7/4".
	TimeSignature int64 `json:"time_signature"`
	// The confidence, from 0.0 to 1.0, of the reliability of the `time_signature`.
	// Sections with time signature changes may correspond to low values in this field.
	TimeSignatureConfidence float64 `json:"time_signature_confidence"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence              respjson.Field
		Duration                respjson.Field
		Key                     respjson.Field
		KeyConfidence           respjson.Field
		Loudness                respjson.Field
		Mode                    respjson.Field
		ModeConfidence          respjson.Field
		Start                   respjson.Field
		Tempo                   respjson.Field
		TempoConfidence         respjson.Field
		TimeSignature           respjson.Field
		TimeSignatureConfidence respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioAnalysisGetResponseSection) RawJSON() string { return r.JSON.raw }
func (r *AudioAnalysisGetResponseSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioAnalysisGetResponseSegment struct {
	// The confidence, from 0.0 to 1.0, of the reliability of the segmentation.
	// Segments of the song which are difficult to logically segment (e.g: noise) may
	// correspond to low values in this field.
	Confidence float64 `json:"confidence"`
	// The duration (in seconds) of the segment.
	Duration float64 `json:"duration"`
	// The offset loudness of the segment in decibels (dB). This value should be
	// equivalent to the loudness_start of the following segment.
	LoudnessEnd float64 `json:"loudness_end"`
	// The peak loudness of the segment in decibels (dB). Combined with
	// `loudness_start` and `loudness_max_time`, these components can be used to
	// describe the "attack" of the segment.
	LoudnessMax float64 `json:"loudness_max"`
	// The segment-relative offset of the segment peak loudness in seconds. Combined
	// with `loudness_start` and `loudness_max`, these components can be used to
	// desctibe the "attack" of the segment.
	LoudnessMaxTime float64 `json:"loudness_max_time"`
	// The onset loudness of the segment in decibels (dB). Combined with `loudness_max`
	// and `loudness_max_time`, these components can be used to describe the "attack"
	// of the segment.
	LoudnessStart float64 `json:"loudness_start"`
	// Pitch content is given by a “chroma” vector, corresponding to the 12 pitch
	// classes C, C#, D to B, with values ranging from 0 to 1 that describe the
	// relative dominance of every pitch in the chromatic scale. For example a C Major
	// chord would likely be represented by large values of C, E and G (i.e. classes 0,
	// 4, and 7).
	//
	// Vectors are normalized to 1 by their strongest dimension, therefore noisy sounds
	// are likely represented by values that are all close to 1, while pure tones are
	// described by one value at 1 (the pitch) and others near 0. As can be seen below,
	// the 12 vector indices are a combination of low-power spectrum values at their
	// respective pitch frequencies. ![pitch vector](/assets/audio/Pitch_vector.png)
	Pitches []float64 `json:"pitches"`
	// The starting point (in seconds) of the segment.
	Start float64 `json:"start"`
	// Timbre is the quality of a musical note or sound that distinguishes different
	// types of musical instruments, or voices. It is a complex notion also referred to
	// as sound color, texture, or tone quality, and is derived from the shape of a
	// segment’s spectro-temporal surface, independently of pitch and loudness. The
	// timbre feature is a vector that includes 12 unbounded values roughly centered
	// around 0. Those values are high level abstractions of the spectral surface,
	// ordered by degree of importance.
	//
	// For completeness however, the first dimension represents the average loudness of
	// the segment; second emphasizes brightness; third is more closely correlated to
	// the flatness of a sound; fourth to sounds with a stronger attack; etc. See an
	// image below representing the 12 basis functions (i.e. template segments).
	// ![timbre basis functions](/assets/audio/Timbre_basis_functions.png)
	//
	// The actual timbre of the segment is best described as a linear combination of
	// these 12 basis functions weighted by the coefficient values: timbre = c1 x b1 +
	// c2 x b2 + ... + c12 x b12, where c1 to c12 represent the 12 coefficients and b1
	// to b12 the 12 basis functions as displayed below. Timbre vectors are best used
	// in comparison with each other.
	Timbre []float64 `json:"timbre"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence      respjson.Field
		Duration        respjson.Field
		LoudnessEnd     respjson.Field
		LoudnessMax     respjson.Field
		LoudnessMaxTime respjson.Field
		LoudnessStart   respjson.Field
		Pitches         respjson.Field
		Start           respjson.Field
		Timbre          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioAnalysisGetResponseSegment) RawJSON() string { return r.JSON.raw }
func (r *AudioAnalysisGetResponseSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioAnalysisGetResponseTrack struct {
	// The number of channels used for analysis. If 1, all channels are summed together
	// to mono before analysis.
	AnalysisChannels int64 `json:"analysis_channels"`
	// The sample rate used to decode and analyze this track. May differ from the
	// actual sample rate of this track available on Spotify.
	AnalysisSampleRate int64 `json:"analysis_sample_rate"`
	// A version number for the Echo Nest Musical Fingerprint format used in the
	// codestring field.
	CodeVersion float64 `json:"code_version"`
	// An
	// [Echo Nest Musical Fingerprint (ENMFP)](https://academiccommons.columbia.edu/doi/10.7916/D8Q248M4)
	// codestring for this track.
	Codestring string `json:"codestring"`
	// Length of the track in seconds.
	Duration float64 `json:"duration"`
	// A version number for the EchoPrint format used in the echoprintstring field.
	EchoprintVersion float64 `json:"echoprint_version"`
	// An [EchoPrint](https://github.com/spotify/echoprint-codegen) codestring for this
	// track.
	Echoprintstring string `json:"echoprintstring"`
	// The time, in seconds, at which the track's fade-in period ends. If the track has
	// no fade-in, this will be 0.0.
	EndOfFadeIn float64 `json:"end_of_fade_in"`
	// The key the track is in. Integers map to pitches using standard
	// [Pitch Class notation](https://en.wikipedia.org/wiki/Pitch_class). E.g. 0 = C, 1
	// = C♯/D♭, 2 = D, and so on. If no key was detected, the value is -1.
	Key int64 `json:"key"`
	// The confidence, from 0.0 to 1.0, of the reliability of the `key`.
	KeyConfidence float64 `json:"key_confidence"`
	// The overall loudness of a track in decibels (dB). Loudness values are averaged
	// across the entire track and are useful for comparing relative loudness of
	// tracks. Loudness is the quality of a sound that is the primary psychological
	// correlate of physical strength (amplitude). Values typically range between -60
	// and 0 db.
	Loudness float64 `json:"loudness"`
	// Mode indicates the modality (major or minor) of a track, the type of scale from
	// which its melodic content is derived. Major is represented by 1 and minor is 0.
	Mode int64 `json:"mode"`
	// The confidence, from 0.0 to 1.0, of the reliability of the `mode`.
	ModeConfidence float64 `json:"mode_confidence"`
	// The exact number of audio samples analyzed from this track. See also
	// `analysis_sample_rate`.
	NumSamples int64 `json:"num_samples"`
	// An offset to the start of the region of the track that was analyzed. (As the
	// entire track is analyzed, this should always be 0.)
	OffsetSeconds int64 `json:"offset_seconds"`
	// A version number for the Rhythmstring used in the rhythmstring field.
	RhythmVersion float64 `json:"rhythm_version"`
	// A Rhythmstring for this track. The format of this string is similar to the
	// Synchstring.
	Rhythmstring string `json:"rhythmstring"`
	// This field will always contain the empty string.
	SampleMd5 string `json:"sample_md5"`
	// The time, in seconds, at which the track's fade-out period starts. If the track
	// has no fade-out, this should match the track's length.
	StartOfFadeOut float64 `json:"start_of_fade_out"`
	// A version number for the Synchstring used in the synchstring field.
	SynchVersion float64 `json:"synch_version"`
	// A [Synchstring](https://github.com/echonest/synchdata) for this track.
	Synchstring string `json:"synchstring"`
	// The overall estimated tempo of a track in beats per minute (BPM). In musical
	// terminology, tempo is the speed or pace of a given piece and derives directly
	// from the average beat duration.
	Tempo float64 `json:"tempo"`
	// The confidence, from 0.0 to 1.0, of the reliability of the `tempo`.
	TempoConfidence float64 `json:"tempo_confidence"`
	// An estimated time signature. The time signature (meter) is a notational
	// convention to specify how many beats are in each bar (or measure). The time
	// signature ranges from 3 to 7 indicating time signatures of "3/4", to "7/4".
	TimeSignature int64 `json:"time_signature"`
	// The confidence, from 0.0 to 1.0, of the reliability of the `time_signature`.
	TimeSignatureConfidence float64 `json:"time_signature_confidence"`
	// The length of the region of the track was analyzed, if a subset of the track was
	// analyzed. (As the entire track is analyzed, this should always be 0.)
	WindowSeconds int64 `json:"window_seconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnalysisChannels        respjson.Field
		AnalysisSampleRate      respjson.Field
		CodeVersion             respjson.Field
		Codestring              respjson.Field
		Duration                respjson.Field
		EchoprintVersion        respjson.Field
		Echoprintstring         respjson.Field
		EndOfFadeIn             respjson.Field
		Key                     respjson.Field
		KeyConfidence           respjson.Field
		Loudness                respjson.Field
		Mode                    respjson.Field
		ModeConfidence          respjson.Field
		NumSamples              respjson.Field
		OffsetSeconds           respjson.Field
		RhythmVersion           respjson.Field
		Rhythmstring            respjson.Field
		SampleMd5               respjson.Field
		StartOfFadeOut          respjson.Field
		SynchVersion            respjson.Field
		Synchstring             respjson.Field
		Tempo                   respjson.Field
		TempoConfidence         respjson.Field
		TimeSignature           respjson.Field
		TimeSignatureConfidence respjson.Field
		WindowSeconds           respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioAnalysisGetResponseTrack) RawJSON() string { return r.JSON.raw }
func (r *AudioAnalysisGetResponseTrack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
