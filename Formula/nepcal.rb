class Nepcal < Formula
  desc "Equivalent of Linux's cal, for Nepali dates"
  homepage "https://github.com/nepcal/nepcal"
  url "https://github.com/srishanbhattarai/nepcal/releases/download/v0.4.0/nepcal_0.4.0_darwin_amd64.tar.gz"
  version "0.4.0"
  sha256 "39bc4a91c44623a947c0fc32d1292ab8f17fd1637b1db107abb1f00976716d9d"

  def install
    bin.install "nepcal"
  end
end
