// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"github.com/cjavdev/spotted-go/internal/apierror"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/cjavdev/spotted-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type AlbumRestrictionObject = shared.AlbumRestrictionObject

// The reason for the restriction. Albums may be restricted if the content is not
// available in a given market, to the user's subscription type, or when the user's
// account is set to not play explicit content. Additional reasons may be added in
// the future.
//
// This is an alias to an internal type.
type AlbumRestrictionObjectReason = shared.AlbumRestrictionObjectReason

// Equals "market"
const AlbumRestrictionObjectReasonMarket = shared.AlbumRestrictionObjectReasonMarket

// Equals "product"
const AlbumRestrictionObjectReasonProduct = shared.AlbumRestrictionObjectReasonProduct

// Equals "explicit"
const AlbumRestrictionObjectReasonExplicit = shared.AlbumRestrictionObjectReasonExplicit

// This is an alias to an internal type.
type ArtistObject = shared.ArtistObject

// The object type.
//
// This is an alias to an internal type.
type ArtistObjectType = shared.ArtistObjectType

// Equals "artist"
const ArtistObjectTypeArtist = shared.ArtistObjectTypeArtist

// This is an alias to an internal type.
type AudiobookBase = shared.AudiobookBase

// This is an alias to an internal type.
type AuthorObject = shared.AuthorObject

// This is an alias to an internal type.
type ChapterRestrictionObject = shared.ChapterRestrictionObject

// This is an alias to an internal type.
type CopyrightObject = shared.CopyrightObject

// This is an alias to an internal type.
type EpisodeObject = shared.EpisodeObject

// The precision with which `release_date` value is known.
//
// This is an alias to an internal type.
type EpisodeObjectReleaseDatePrecision = shared.EpisodeObjectReleaseDatePrecision

// Equals "year"
const EpisodeObjectReleaseDatePrecisionYear = shared.EpisodeObjectReleaseDatePrecisionYear

// Equals "month"
const EpisodeObjectReleaseDatePrecisionMonth = shared.EpisodeObjectReleaseDatePrecisionMonth

// Equals "day"
const EpisodeObjectReleaseDatePrecisionDay = shared.EpisodeObjectReleaseDatePrecisionDay

// This is an alias to an internal type.
type EpisodeRestrictionObject = shared.EpisodeRestrictionObject

// This is an alias to an internal type.
type ExternalIDObject = shared.ExternalIDObject

// This is an alias to an internal type.
type ExternalURLObject = shared.ExternalURLObject

// This is an alias to an internal type.
type FollowersObject = shared.FollowersObject

// This is an alias to an internal type.
type ImageObject = shared.ImageObject

// This is an alias to an internal type.
type LinkedTrackObject = shared.LinkedTrackObject

// This is an alias to an internal type.
type NarratorObject = shared.NarratorObject

// This is an alias to an internal type.
type PagingPlaylistObject = shared.PagingPlaylistObject

// This is an alias to an internal type.
type PlaylistTrackObject = shared.PlaylistTrackObject

// Information about the track or episode.
//
// This is an alias to an internal type.
type PlaylistTrackObjectTrackUnion = shared.PlaylistTrackObjectTrackUnion

// This is an alias to an internal type.
type PlaylistTracksRefObject = shared.PlaylistTracksRefObject

// This is an alias to an internal type.
type PlaylistUserObject = shared.PlaylistUserObject

// The object type.
//
// This is an alias to an internal type.
type PlaylistUserObjectType = shared.PlaylistUserObjectType

// Equals "user"
const PlaylistUserObjectTypeUser = shared.PlaylistUserObjectTypeUser

// This is an alias to an internal type.
type ResumePointObject = shared.ResumePointObject

// This is an alias to an internal type.
type ShowBase = shared.ShowBase

// This is an alias to an internal type.
type SimplifiedArtistObject = shared.SimplifiedArtistObject

// The object type.
//
// This is an alias to an internal type.
type SimplifiedArtistObjectType = shared.SimplifiedArtistObjectType

// Equals "artist"
const SimplifiedArtistObjectTypeArtist = shared.SimplifiedArtistObjectTypeArtist

// This is an alias to an internal type.
type SimplifiedEpisodeObject = shared.SimplifiedEpisodeObject

// The precision with which `release_date` value is known.
//
// This is an alias to an internal type.
type SimplifiedEpisodeObjectReleaseDatePrecision = shared.SimplifiedEpisodeObjectReleaseDatePrecision

// Equals "year"
const SimplifiedEpisodeObjectReleaseDatePrecisionYear = shared.SimplifiedEpisodeObjectReleaseDatePrecisionYear

// Equals "month"
const SimplifiedEpisodeObjectReleaseDatePrecisionMonth = shared.SimplifiedEpisodeObjectReleaseDatePrecisionMonth

// Equals "day"
const SimplifiedEpisodeObjectReleaseDatePrecisionDay = shared.SimplifiedEpisodeObjectReleaseDatePrecisionDay

// This is an alias to an internal type.
type SimplifiedPlaylistObject = shared.SimplifiedPlaylistObject

// The user who owns the playlist
//
// This is an alias to an internal type.
type SimplifiedPlaylistObjectOwner = shared.SimplifiedPlaylistObjectOwner

// This is an alias to an internal type.
type SimplifiedTrackObject = shared.SimplifiedTrackObject

// This is an alias to an internal type.
type TrackObject = shared.TrackObject

// The album on which the track appears. The album object includes a link in `href`
// to full information about the album.
//
// This is an alias to an internal type.
type TrackObjectAlbum = shared.TrackObjectAlbum

// The object type: "track".
//
// This is an alias to an internal type.
type TrackObjectType = shared.TrackObjectType

// Equals "track"
const TrackObjectTypeTrack = shared.TrackObjectTypeTrack

// This is an alias to an internal type.
type TrackRestrictionObject = shared.TrackRestrictionObject
