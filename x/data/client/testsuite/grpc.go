package testsuite

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/regen-network/regen-ledger/x/data"
)

func (s *IntegrationTestSuite) TestQueryByIRI() {
	val := s.network.Validators[0]

	iri := "regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf"

	testCases := []struct {
		name   string
		url    string
		expErr bool
		errMsg string
	}{
		{
			"invalid IRI",
			fmt.Sprintf("%s/regen/data/v1/by-iri/%s", val.APIAddress, "foo"),
			true,
			"key not found",
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/data/v1/by-iri/%s", val.APIAddress, iri),
			false,
			"",
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var entry data.QueryByIRIResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &entry)

			if tc.expErr {
				require.Error(err)
			} else {
				require.NoError(err)
				require.NotNil(entry.Entry)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryByAttestor() {
	val := s.network.Validators[0]

	acc1, err := val.ClientCtx.Keyring.Key("acc1")
	s.Require().NoError(err)

	addr := acc1.GetAddress().String()

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		errMsg   string
		expItems int
	}{
		{
			"invalid attestor",
			fmt.Sprintf("%s/regen/data/v1/by-attestor/%s", val.APIAddress, "foo"),
			true,
			"invalid bech32 string",
			0,
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/data/v1/by-attestor/%s", val.APIAddress, addr),
			false,
			"",
			2,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/regen/data/v1/by-attestor/%s?pagination.limit=1", val.APIAddress, addr),
			false,
			"",
			1,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var entries data.QueryByAttestorResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &entries)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(entries.Entries)
				require.Len(entries.Entries, tc.expItems)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAttestors() {
	val := s.network.Validators[0]

	iri := "regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf"

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		errMsg   string
		expItems int
	}{
		{
			"invalid attestor",
			fmt.Sprintf("%s/regen/data/v1/attestors/%s", val.APIAddress, "foo"),
			true,
			"key not found",
			0,
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/data/v1/attestors/%s", val.APIAddress, iri),
			false,
			"",
			2,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/regen/data/v1/attestors/%s?pagination.limit=1", val.APIAddress, iri),
			false,
			"",
			1,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var attestors data.QueryAttestorsResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &attestors)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(attestors.Attestors)
				require.Len(attestors.Attestors, tc.expItems)
			}
		})
	}
}
